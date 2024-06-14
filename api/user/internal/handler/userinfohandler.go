package handler

import (
	"net/http"

	"go-micro/api/user/internal/logic"
	"go-micro/api/user/internal/svc"
	"go-micro/utils"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUserInfoLogic(r.Context(), svcCtx)
		resp, stat, err := l.UserInfo()
		if stat < 0 {
			utils.ErrorJson(w, stat, err.Error())
		} else {
			utils.SuccessJson(w, resp)
		}
	}
}
