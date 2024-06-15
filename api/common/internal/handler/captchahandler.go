package handler

import (
	"net/http"

	"go-micro/api/common/internal/logic"
	"go-micro/api/common/internal/svc"
	"go-micro/utils"
)

func CaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewCaptchaLogic(r.Context(), svcCtx)
		resp, stat, err := l.Captcha()
		if stat < 0 {
			utils.ErrorJson(w, stat, err.Error())
		} else {
			utils.SuccessJson(w, "", resp)
		}
	}
}
