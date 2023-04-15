package sys_user_model

import (
	"context"
	"go-zero-micro/pkg/models"
)

func FindOneById(ctx context.Context, id int64) (*SysUser, error) {
	var user = &SysUser{}
	err := models.DB.First(user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
