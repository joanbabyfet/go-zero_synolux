package handler

import (
	"net/http"

	"go-micro/api/user/internal/logic"
	"go-micro/api/user/internal/svc"
	"go-micro/utils"
)

func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLogoutLogic(r.Context(), svcCtx)
		stat, err := l.Logout()
		if stat < 0 {
			utils.ErrorJson(w, stat, err.Error())
		} else {
			utils.SuccessJson(w, nil)
		}
	}
}
