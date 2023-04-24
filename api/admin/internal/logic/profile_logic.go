package logic

import (
	"context"
	"go-zero-micro/api/admin/internal/response"
	"go-zero-micro/api/admin/internal/svc"
	"go-zero-micro/api/admin/internal/types"
	"go-zero-micro/common/errorx"
	"go-zero-micro/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProfileLogic {
	return &ProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProfileLogic) Profile() (resp *types.UserReply, err error) {
	userId := GetUserId(l.ctx)
	sysUser, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
		Id: userId,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("用户不存在!")
	}
	return response.OfUserReply(sysUser), nil
}
