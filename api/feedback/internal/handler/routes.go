// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-micro/api/feedback/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/feedback",
				Handler: SaveFeedbackHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}