package sys_user_model

import (
	"context"
	"go-zero-micro/common/models"
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

func FindOneById(ctx context.Context, id int64) (*SysUser, error) {
	var u = &SysUser{}
	err := models.DB.First(u, id).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func FindOneByLogin(username string, password string) (*SysUser, error) {
	var u *SysUser
	err := models.DB.Where(&SysUser{
		Username: username,
		Password: password,
	}).Find(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}
