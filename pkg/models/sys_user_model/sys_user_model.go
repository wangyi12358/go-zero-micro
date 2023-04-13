package sys_user_model

import (
	"context"
	"go-zero-micro/pkg/models"
	"go-zero-micro/pkg/models/gen"
	"go-zero-micro/pkg/models/model"
)

func FindOneById(ctx context.Context, id int64) (*model.SysUser, error) {
	query := gen.Use(models.DB).SysUser
	sysUser, err := query.WithContext(ctx).Where(query.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return sysUser, nil
}
