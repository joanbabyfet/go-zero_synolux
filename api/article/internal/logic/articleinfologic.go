package logic

import (
	"context"
	"errors"

	"go-micro/api/article/internal/svc"
	"go-micro/api/article/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleInfoLogic {
	return &ArticleInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleInfoLogic) ArticleInfo(req *types.IdRequest) (*types.ArticleInfo, int, error) {
	stat := 1
	//获取文章详情
	data, err := l.svcCtx.ArticleModel.FindOne(l.ctx, req.Id)
	if err != nil {
		return nil, -2, errors.New("文章不存在")
	}
	info := &types.ArticleInfo{
		Id:      data.Id,
		Catid:   data.Catid,
		Title:   data.Title,
		Info:    data.Info,
		Content: data.Content.String,
		Author:  data.Author,
		Status:  data.Status,
	}
	return info, stat, nil
}
