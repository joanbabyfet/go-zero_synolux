package svc

import (
	"go-micro/api/user/internal/config"
	"go-micro/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	// 定義 Model 結構體
	UserModel model.UserModel
	//客户端ip
	ClientIP string
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 数据库连接
	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,
		// 把 UserModel 对象 new 出來
		UserModel: model.NewUserModel(sqlConn),
	}
}
