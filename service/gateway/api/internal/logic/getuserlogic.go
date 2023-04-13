package logic

import (
	"context"
	user2 "go-zero-micro/service/user/rpc/types/user"
	"strconv"

	"go-zero-micro/service/gateway/api/internal/svc"
	"go-zero-micro/service/gateway/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetProductReq) (resp *types.UserReply, err error) {
	id, _ := strconv.ParseInt(req.Id, 10, 64)
	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user2.IdRequest{
		Id: id,
	})
	return &types.UserReply{
		Nickname: user.Nickname,
		Gender:   int16(user.Gender),
	}, nil
}
