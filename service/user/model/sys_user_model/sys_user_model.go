package sys_user_model

import (
	"context"
	"fmt"
	"go-zero-micro/core/model"
	"go-zero-micro/service/user/rpc/types/user"
)

func OfUserResponse(sysUser *SysUser) *user.UserResponse {
	if sysUser == nil {
		return nil
	}
	return &user.UserResponse{
		Id:       sysUser.ID,
		Nickname: sysUser.Nickname,
		Gender:   int32(sysUser.Gender),
	}
}

func OfUserResponses(users []SysUser) []*user.UserResponse {
	var list = make([]*user.UserResponse, len(users))
	for i, u := range users {
		list[i] = OfUserResponse(&u)
	}
	return list
}

func FindOneById(ctx context.Context, id int64) (*SysUser, error) {
	var u = &SysUser{}
	err := model.DB.First(u, id).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func FindOneByLogin(username string, password string) (*SysUser, error) {
	u := &SysUser{}
	err := model.DB.Where(&SysUser{
		Username: username,
		Password: password,
	}).First(u).Error
	if err != nil {
		fmt.Printf("sql error: %s\n", err.Error())
		return nil, err
	}
	return u, nil
}

func FindUserPage(in *user.UserPageReq) (res *user.UserPageRes, err error) {
	var users []SysUser
	var total int64
	offset := in.PageSize * (in.Current - 1)
	err = model.DB.Model(&SysUser{}).Offset(int(offset)).Limit(int(in.PageSize)).Order("id desc").Find(&users).Error
	if err != nil {
		return nil, err
	}
	err = model.DB.Model(&SysUser{}).Count(&total).Error
	if err != nil {
		return nil, err
	}
	return &user.UserPageRes{
		Total: int32(total),
		List:  OfUserResponses(users),
	}, nil
}
