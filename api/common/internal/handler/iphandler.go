package handler

import (
	"net/http"

	"go-micro/api/common/internal/logic"
	"go-micro/api/common/internal/svc"
	"go-micro/utils"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func IpHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//获取客户端ip
		svcCtx.ClientIP = httpx.GetRemoteAddr(r)
		l := logic.NewIpLogic(r.Context(), svcCtx)
		resp, stat, err := l.Ip()
		if stat < 0 {
			utils.ErrorJson(w, stat, err.Error())
		} else {
			utils.SuccessJson(w, "", resp)
		}
	}
}
