package logic

import (
	"context"

	"go-micro/api/article/internal/svc"
	"go-micro/api/article/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type HomeArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeArticleLogic {
	return &HomeArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeArticleLogic) HomeArticle(req *types.SearchRequest) (res map[string]interface{}, stat int, err error) {
	stat = 1
	//获取列表
	articleList, err := l.svcCtx.ArticleModel.FindAll(l.ctx, req.Title, req.Limit)
	if err != nil {
		return nil, -2, err
	}
	var articleListTmp []*types.ArticleInfo
	_ = copier.Copy(&articleListTmp, articleList)

	//组装数据
	res = make(map[string]interface{}) //创建1个空集合
	res["list"] = articleListTmp

	return res, stat, nil
}
