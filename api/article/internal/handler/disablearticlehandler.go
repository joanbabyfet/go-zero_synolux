package handler

import (
	"net/http"

	"go-micro/api/article/internal/logic"
	"go-micro/api/article/internal/svc"
	"go-micro/api/article/internal/types"
	"go-micro/utils"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func DisableArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdRequest
		if err := httpx.Parse(r, &req); err != nil {
			utils.ErrorJson(w, -1, err.Error())
			return
		}

		l := logic.NewDisableArticleLogic(r.Context(), svcCtx)
		stat, err := l.DisableArticle(&req)
		if stat < 0 {
			utils.ErrorJson(w, stat, err.Error())
		} else {
			utils.SuccessJson(w, nil)
		}
	}
}
