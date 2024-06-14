package handler

import (
	"net/http"

	"go-micro/api/user/internal/logic"
	"go-micro/api/user/internal/svc"
	"go-micro/api/user/internal/types"
	"go-micro/utils"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			utils.ErrorJson(w, -1, err.Error())
			return
		}

		//获取客户端ip
		svcCtx.ClientIP = httpx.GetRemoteAddr(r)
		l := logic.NewCreateUserLogic(r.Context(), svcCtx)
		stat, err := l.CreateUser(&req)
		if stat < 0 {
			utils.ErrorJson(w, stat, err.Error())
		} else {
			utils.SuccessJson(w, nil)
		}
	}
}
