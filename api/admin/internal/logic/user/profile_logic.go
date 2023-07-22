package user

import (
	"context"
	"encoding/json"
	"fmt"
	"go-zero-micro/api/admin/internal/logic/admin"
	"go-zero-micro/api/admin/internal/response"
	"go-zero-micro/api/admin/internal/svc"
	"go-zero-micro/api/admin/internal/types"
	"go-zero-micro/common/errorx"
	"go-zero-micro/core/redis"
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

func setRedisProfile(l *ProfileLogic, userReply *types.UserReply) {
	userJSON, err := json.Marshal(userReply)
	if err != nil {
		panic(err)
	}
	key := fmt.Sprintf("user-%d", userReply.Id)
	redis.DB.Set(l.ctx, key, userJSON, 0)
}

func getRedisProfile(l *ProfileLogic, userId int64) *types.UserReply {
	var userReply types.UserReply
	key := fmt.Sprintf("user-%d", userId)
	userJSON, err := redis.DB.Get(l.ctx, key).Result()
	if err != nil {
		return nil
	}
	err = json.Unmarshal([]byte(userJSON), &userReply)
	if err != nil {
		return nil
	}
	return &userReply
}

func (l *ProfileLogic) Profile() (resp *types.UserReply, err error) {
	userId := admin.GetUserId(l.ctx)
	if resp = getRedisProfile(l, userId); resp != nil {
		return resp, nil
	}
	sysUser, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
		Id: userId,
	})
	get := redis.DB.Get(l.ctx, fmt.Sprintf("user-%d", userId))
	get.Val()
	if err != nil {
		return nil, errorx.NewDefaultError("用户不存在!")
	}
	resp = response.OfUserReply(sysUser)
	setRedisProfile(l, resp)
	return resp, nil
}
