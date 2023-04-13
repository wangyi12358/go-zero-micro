package logic

import (
	"context"

	"go-zero-micro/service/user/rpc/internal/svc"
	"go-zero-micro/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserLogic {
	return &CheckUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserLogic) CheckUser(in *user.CheckUserRequest) (*user.CheckUserResponse, error) {
	// todo: add your logic here and delete this line

	return &user.CheckUserResponse{}, nil
}
