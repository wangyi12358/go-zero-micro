package logic

import (
	"context"
	"go-zero-micro/service/user/model/sys_user_model"

	"go-zero-micro/service/user/rpc/internal/svc"
	"go-zero-micro/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPageLogic {
	return &UserPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserPageLogic) UserPage(in *user.UserPageReq) (*user.UserPageRes, error) {
	res, err := sys_user_model.FindUserPage(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
