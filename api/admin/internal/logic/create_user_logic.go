package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-micro/api/admin/internal/svc"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser() error {
	// todo: add your logic here and delete this line

	return nil
}
