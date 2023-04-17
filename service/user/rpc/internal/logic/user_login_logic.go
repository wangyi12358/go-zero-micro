package logic

import (
	"context"
	"go-zero-micro/service/user/model/sys_user_model"

	"go-zero-micro/service/user/rpc/internal/svc"
	"go-zero-micro/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserLoginLogic) UserLogin(in *user.CheckUserRequest) (*user.UserResponse, error) {
	sysUser, err := sys_user_model.FindOneByLogin(in.Username, in.Password)
	if err == nil {
		return nil, err
	}

	return sys_user_model.OfUserResponse(sysUser), nil
}
