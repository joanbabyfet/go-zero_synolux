package handler

import (
	"net/http"

	"go-micro/api/common/internal/logic"
	"go-micro/api/common/internal/svc"
	"go-micro/utils"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPingLogic(r.Context(), svcCtx)
		stat, err := l.Ping()
		if stat < 0 {
			utils.ErrorJson(w, stat, err.Error())
		} else {
			utils.SuccessJson(w, "pong", nil)
		}
	}
}
