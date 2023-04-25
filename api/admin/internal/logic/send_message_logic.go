package logic

import (
	"context"
	"github.com/Shopify/sarama"
	"go-zero-micro/common/errorx"
	"go-zero-micro/core/kafka"

	"go-zero-micro/api/admin/internal/svc"
	"go-zero-micro/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMessageLogic) SendMessage(req *types.SendMessageReq) error {
	_, _, err := kafka.Producer.SendMessage(&sarama.ProducerMessage{
		Topic: req.Topic,
		Value: sarama.StringEncoder(req.Message),
	})
	if err != nil {
		return errorx.NewDefaultError("kafka数据写入失败!")
	}
	return nil
}
