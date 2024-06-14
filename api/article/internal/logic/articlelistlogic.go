package logic

import (
	"context"

	"go-micro/api/article/internal/svc"
	"go-micro/api/article/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleListLogic {
	return &ArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleListLogic) ArticleList(req *types.SearchRequest) (res map[string]interface{}, stat int, err error) {
	stat = 1
	//获取分页列表
	articleList, total, err := l.svcCtx.ArticleModel.PageList(l.ctx, req.Page, req.PageSize, req.Title)
	if err != nil {
		return nil, -2, err
	}
	var articleListTmp []*types.ArticleInfo
	_ = copier.Copy(&articleListTmp, articleList)

	//组装数据
	res = make(map[string]interface{}) //创建1个空集合
	res["total"] = total
	res["list"] = articleListTmp

	return res, stat, nil
}
