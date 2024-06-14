package logic

import (
	"context"
	"database/sql"
	"errors"

	"go-micro/api/article/internal/svc"
	"go-micro/api/article/internal/types"
	"go-micro/model"
	"go-micro/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveArticleLogic {
	return &SaveArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveArticleLogic) SaveArticle(req *types.SaveRequest) (stat int, err error) {
	stat = 1
	if req.Id > 0 {
		//修改
		data, _ := l.svcCtx.ArticleModel.FindOne(l.ctx, req.Id)
		data.Catid = req.Catid
		data.Title = req.Title
		data.Info = req.Info
		data.Content = sql.NullString{
			String: req.Content,
			Valid:  true,
		}
		data.Author = req.Author
		data.Status = req.Status
		data.UpdateTime = int64(utils.Timestamp())
		data.UpdateUser = "1" //修改人
		err = l.svcCtx.ArticleModel.Update(l.ctx, data)
	} else {
		//添加
		_, err = l.svcCtx.ArticleModel.Insert(l.ctx, &model.Article{
			Catid: req.Catid,
			Title: req.Title,
			Info:  req.Info,
			Content: sql.NullString{
				String: req.Content,
				Valid:  true,
			},
			Author:     req.Author,
			Status:     1,
			CreateTime: int64(utils.Timestamp()), //添加时间
			CreateUser: "1",                      //添加人
		})
	}
	if err != nil {
		//错误信息写在这
		return -2, errors.New("文章保存失败")
	}
	return stat, nil
}
