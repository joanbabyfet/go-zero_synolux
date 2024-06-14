package handler

import (
	"net/http"

	"go-micro/api/feedback/internal/logic"
	"go-micro/api/feedback/internal/svc"
	"go-micro/api/feedback/internal/types"
	"go-micro/utils"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SaveFeedbackHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SaveRequest
		if err := httpx.Parse(r, &req); err != nil {
			utils.ErrorJson(w, -1, err.Error())
			return
		}

		l := logic.NewSaveFeedbackLogic(r.Context(), svcCtx)
		stat, err := l.SaveFeedback(&req)
		if stat < 0 {
			utils.ErrorJson(w, stat, err.Error())
		} else {
			utils.SuccessJson(w, nil)
		}
	}
}
