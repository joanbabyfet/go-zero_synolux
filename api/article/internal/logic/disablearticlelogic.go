package logic

import (
	"context"
	"errors"

	"go-micro/api/article/internal/svc"
	"go-micro/api/article/internal/types"
	"go-micro/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type DisableArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDisableArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DisableArticleLogic {
	return &DisableArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DisableArticleLogic) DisableArticle(req *types.IdRequest) (int, error) {
	stat := 1
	//获取文章详情
	data, _ := l.svcCtx.ArticleModel.FindOne(l.ctx, req.Id)
	data.Status = 0
	data.UpdateTime = int64(utils.Timestamp())
	data.UpdateUser = "1" //修改人
	//更新
	err := l.svcCtx.ArticleModel.Update(l.ctx, data)
	if err != nil {
		//错误信息写在这
		return -2, errors.New("文章禁用失败")
	}
	return stat, nil
}
