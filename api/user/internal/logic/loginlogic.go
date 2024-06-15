package logic

import (
	"context"
	"errors"
	"strings"

	"go-micro/api/user/internal/svc"
	"go-micro/api/user/internal/types"
	"go-micro/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (*types.LoginResponse, int, error) {
	stat := 1
	auth := l.svcCtx.Config.Auth

	//检测验证码
	if !utils.Store.Verify(req.Key, req.Code, true) {
		return nil, -2, errors.New("验证码错误")
	}

	//获取用户详情
	data, err := l.svcCtx.UserModel.FindByUsername(l.ctx, req.Username)
	if err != nil {
		return nil, -3, errors.New("用户名或密码无效")
	}
	if !utils.PasswordVerify(strings.Trim(req.Password, " "), data.Password) {
		return nil, -4, errors.New("用户名或密码无效")
	}
	data.LoginIp = l.svcCtx.ClientIP //登录ip
	data.LoginTime = int64(utils.Timestamp())
	err = l.svcCtx.UserModel.Update(l.ctx, data)
	if err != nil {
		return nil, -5, errors.New("登录异常")
	}

	// 生成token
	token, err := utils.GetToken(utils.JwtPayLoad{
		UserID:   data.Id,
		Username: data.Username,
	}, auth.AccessSecret, auth.AccessExpire)
	if err != nil {
		return nil, -6, errors.New("生成token失败")
	}
	//测试解析token
	// parse_token, _ := utils.ParseToken(token, auth.AccessSecret)
	// fmt.Println(parse_token)

	//组装返回数据
	info := &types.LoginResponse{
		Id:        data.Id,
		Username:  data.Username,
		Realname:  data.Realname,
		Email:     data.Email,
		PhoneCode: data.PhoneCode,
		Phone:     data.Phone,
		Avatar:    utils.DisplayImg(l.svcCtx.Config.FileUrl, data.Avatar),
		Language:  data.Language,
		Token:     token,
	}
	return info, stat, nil
}
