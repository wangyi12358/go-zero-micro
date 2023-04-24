package response

import (
	"go-zero-micro/api/admin/internal/types"
	"go-zero-micro/service/user/rpc/types/user"
)

func OfUserReply(userRes *user.UserResponse) *types.UserReply {
	return &types.UserReply{
		Id:       userRes.Id,
		Nickname: userRes.Nickname,
		Gender:   userRes.Gender,
	}
}

func OfUserPageReply(res *user.UserPageRes) *types.UserPageReply {
	var list = make([]types.UserReply, len(res.List))
	for i, u := range res.List {
		list[i] = *OfUserReply(u)
	}
	return &types.UserPageReply{
		List:  list,
		Total: res.Total,
	}
}
