package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-micro/api/admin/internal/svc"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList() error {
	// todo: add your logic here and delete this line

	return nil
}
