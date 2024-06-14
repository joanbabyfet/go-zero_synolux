package logic

import (
	"context"
	"errors"

	"go-micro/api/user/internal/svc"
	"go-micro/api/user/internal/types"
	"go-micro/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (*types.UserInfo, int, error) {
	stat := 1
	//拿到token里用户id
	uid := l.ctx.Value("user_id").(string)
	//获取用户详情
	data, err := l.svcCtx.UserModel.FindOne(l.ctx, uid)
	if err != nil {
		return nil, -2, errors.New("用户不存在")
	}
	info := &types.UserInfo{
		Id:        data.Id,
		Username:  data.Username,
		Realname:  data.Realname,
		Email:     data.Email,
		PhoneCode: data.PhoneCode,
		Phone:     data.Phone,
		Avatar:    utils.DisplayImg(l.svcCtx.Config.FileUrl, data.Avatar),
		Language:  data.Language,
	}
	return info, stat, nil
}
