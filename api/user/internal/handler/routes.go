// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-micro/api/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: CreateUserHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/get_userinfo",
				Handler: UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/logout",
				Handler: LogoutHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/profile",
				Handler: UpdateUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/set_password",
				Handler: SetPasswordHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)
}
