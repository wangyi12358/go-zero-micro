package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-micro/service/user/model/sys_user_model"
	"go-zero-micro/service/user/rpc/internal/svc"
	"go-zero-micro/service/user/rpc/types/user"
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
	sysUser, err := sys_user_model.FindOneById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &user.UserResponse{
		Id:       in.Id,
		Nickname: sysUser.Nickname,
		Gender:   int32(sysUser.Gender),
	}, nil
}
