package logic

import (
	"context"
	"database/sql"
	"errors"

	"go-micro/api/feedback/internal/svc"
	"go-micro/api/feedback/internal/types"
	"go-micro/model"
	"go-micro/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveFeedbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveFeedbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveFeedbackLogic {
	return &SaveFeedbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveFeedbackLogic) SaveFeedback(req *types.SaveRequest) (stat int, err error) {
	stat = 1
	//添加
	_, err = l.svcCtx.FeedbackModel.Insert(l.ctx, &model.Feedback{
		Name:   req.Name,
		Mobile: req.Mobile,
		Email:  req.Email,
		Content: sql.NullString{
			String: req.Content,
			Valid:  true,
		},
		CreateTime: int64(utils.Timestamp()), //添加时间
		CreateUser: "1",                      //添加人
	})
	if err != nil {
		//错误信息写在这
		return -2, errors.New("反馈保存失败")
	}
	return stat, nil
}
