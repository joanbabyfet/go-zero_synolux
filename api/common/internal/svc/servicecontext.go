package svc

import (
	"go-micro/api/common/internal/config"
)

type ServiceContext struct {
	Config config.Config
	//客户端ip
	ClientIP string
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
