package handler

import (
	"net/http"

	"go-micro/api/common/internal/logic"
	"go-micro/api/common/internal/svc"
	"go-micro/utils"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var req types.UploadRequest
		// if err := httpx.Parse(r, &req); err != nil {
		// 	utils.ErrorJson(w, -1, err.Error())
		// 	return
		// }

		l := logic.NewUploadLogic(r.Context(), svcCtx)
		resp, stat, err := l.Upload(r)
		if stat < 0 {
			utils.ErrorJson(w, stat, err.Error())
		} else {
			utils.SuccessJson(w, "", resp)
		}
	}
}
