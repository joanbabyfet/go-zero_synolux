package logic

import (
	"context"

	"go-micro/api/common/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type IpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IpLogic {
	return &IpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IpLogic) Ip() (res map[string]interface{}, stat int, err error) {
	stat = 1

	//组装数据
	resp := make(map[string]interface{}) //创建1个空集合
	resp["ip"] = l.svcCtx.ClientIP
	return resp, stat, nil
}
