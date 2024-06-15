package logic

import (
	"context"
	"errors"

	"go-micro/api/common/internal/svc"
	"go-micro/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaLogic {
	return &CaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaLogic) Captcha() (res map[string]interface{}, stat int, err error) {
	id, b64s, _, err := utils.GetCaptcha()
	if err != nil {
		return nil, -2, errors.New("生成验证码错误")
	}

	//组装数据
	resp := make(map[string]interface{}) //创建1个空集合
	resp["key"] = id
	resp["img"] = b64s
	return resp, stat, nil
}
