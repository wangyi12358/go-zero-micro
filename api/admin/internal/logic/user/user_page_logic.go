package user

import (
	"context"
	"go-zero-micro/api/admin/internal/response"
	"go-zero-micro/service/user/rpc/types/user"

	"go-zero-micro/api/admin/internal/svc"
	"go-zero-micro/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserPageLogic {
	return &UserPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserPageLogic) UserPage(req *types.UserPageReq) (resp *types.UserPageReply, err error) {
	res, err := l.svcCtx.UserRpc.UserPage(l.ctx, &user.UserPageReq{
		Current:  int32(req.Current),
		PageSize: int32(req.PageSize),
	})
	if err != nil {
		return nil, err
	}
	return response.OfUserPageReply(res), nil
}
