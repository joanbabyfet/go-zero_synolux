package svc

import (
	"go-micro/api/feedback/internal/config"
	"go-micro/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	// 定義 Model 結構體
	FeedbackModel model.FeedbackModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 数据库连接
	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,
		// 把 ArticleModel 对象 new 出來
		FeedbackModel: model.NewFeedbackModel(sqlConn),
	}
}
