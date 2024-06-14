package handler

import (
	"net/http"

	"go-micro/api/user/internal/logic"
	"go-micro/api/user/internal/svc"
	"go-micro/api/user/internal/types"
	"go-micro/utils"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SetPasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetPasswordRequest
		if err := httpx.Parse(r, &req); err != nil {
			utils.ErrorJson(w, -1, err.Error())
			return
		}

		l := logic.NewSetPasswordLogic(r.Context(), svcCtx)
		stat, err := l.SetPassword(&req)
		if stat < 0 {
			utils.ErrorJson(w, stat, err.Error())
		} else {
			utils.SuccessJson(w, nil)
		}
	}
}
