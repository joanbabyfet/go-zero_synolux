package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	//数据库配置属性
	DB struct {
		DataSource string
	}
}
