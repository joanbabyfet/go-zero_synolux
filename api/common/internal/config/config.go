package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	FileUrl       string
	UploadMaxSize int
	UploadDir     string
}
