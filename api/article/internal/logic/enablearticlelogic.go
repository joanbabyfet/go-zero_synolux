package logic

import (
	"context"
	"errors"

	"go-micro/api/article/internal/svc"
	"go-micro/api/article/internal/types"
	"go-micro/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnableArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnableArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnableArticleLogic {
	return &EnableArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnableArticleLogic) EnableArticle(req *types.IdRequest) (int, error) {
	stat := 1
	//获取文章详情
	data, _ := l.svcCtx.ArticleModel.FindOne(l.ctx, req.Id)
	data.Status = 1
	data.UpdateTime = int64(utils.Timestamp())
	data.UpdateUser = "1" //修改人
	//更新
	err := l.svcCtx.ArticleModel.Update(l.ctx, data)
	if err != nil {
		//错误信息写在这
		return -2, errors.New("文章启用失败")
	}
	return stat, nil
}
