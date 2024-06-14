package logic

import (
	"context"
	"errors"

	"go-micro/api/user/internal/svc"
	"go-micro/api/user/internal/types"
	"go-micro/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.ProfileRequest) (stat int, err error) {
	stat = 1
	//拿到token里用户id
	uid := l.ctx.Value("user_id").(string)
	//修改
	data, _ := l.svcCtx.UserModel.FindOne(l.ctx, uid)
	data.Realname = req.Realname
	data.Email = req.Email
	data.PhoneCode = req.PhoneCode
	data.Phone = req.Phone
	data.UpdateTime = int64(utils.Timestamp())
	data.UpdateUser = "1" //修改人
	err = l.svcCtx.UserModel.Update(l.ctx, data)
	if err != nil {
		//错误信息写在这
		return -2, errors.New("用户修改失败")
	}
	return stat, nil
}
