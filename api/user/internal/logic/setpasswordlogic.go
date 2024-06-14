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

type SetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetPasswordLogic {
	return &SetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetPasswordLogic) SetPassword(req *types.SetPasswordRequest) (stat int, err error) {
	//检测输入密码是否一致
	if req.RePassword != req.NewPassword {
		return -2, errors.New("确认密码不一样")
	}

	//拿到token里用户id
	uid := l.ctx.Value("user_id").(string)
	//获取用户信息
	data, _ := l.svcCtx.UserModel.FindOne(l.ctx, uid)

	if !utils.PasswordVerify(strings.Trim(req.Password, " "), data.Password) {
		return -3, errors.New("原始密码错误")
	}
	//更新密码
	password, _ := utils.PasswordHash(req.NewPassword)
	data.Password = password
	err = l.svcCtx.UserModel.Update(l.ctx, data)
	if err != nil {
		return -4, errors.New("密码更新失败")
	}
	return stat, nil
}
