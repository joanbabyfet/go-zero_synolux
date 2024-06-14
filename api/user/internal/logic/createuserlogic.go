package logic

import (
	"context"
	"errors"

	"go-micro/api/user/internal/svc"
	"go-micro/api/user/internal/types"
	"go-micro/model"
	"go-micro/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.RegisterRequest) (stat int, err error) {
	// 呼叫前面在 svc 中新定義的 UserModel 下的功能
	// 將請求內容寫入資料庫
	stat = 1
	password, _ := utils.PasswordHash(req.Password)

	_, err = l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Id:         utils.UniqueId(),
		Username:   req.Username,
		Password:   password,
		Realname:   req.Realname,
		Email:      req.Email,
		PhoneCode:  req.PhoneCode,
		Phone:      req.Phone,
		Language:   "cn", //默认中文
		Status:     1,
		RegIp:      l.svcCtx.ClientIP,        //登录ip,
		CreateTime: int64(utils.Timestamp()), //添加时间
		CreateUser: "1",                      //添加人
	})
	if err != nil {
		//错误信息写在这
		return -2, errors.New("用户添加失败")
	}
	return stat, nil
}
