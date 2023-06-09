package logic

import (
	"context"
	"go-zero-micro/service/user/model/sys_user_model"

	"go-zero-micro/service/user/rpc/internal/svc"
	"go-zero-micro/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.IdRequest) (*user.UserResponse, error) {
	sysUser, err := sys_user_model.FindOneById(in.Id)
	if err != nil {
		return nil, err
	}
	return sys_user_model.OfUserResponse(sysUser), nil
}
