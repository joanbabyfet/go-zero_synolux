package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	//数据库配置属性
	DB struct {
		DataSource string
	}

	//jwt验证
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	FileUrl string
}
