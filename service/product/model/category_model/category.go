package category_model

import (
	"go-zero-micro/core/model"
	"go-zero-micro/service/product/rpc/types/product"
)

func OfCategoryResponse(c *Category) *product.CategoryRes {
	return &product.CategoryRes{
		Id:         c.Id,
		Name:       c.Name,
		CreateTime: c.CreateTime.Format("2006-01-02 15:04:05"),
	}
}

func Create(c *Category) error {
	return model.DB.Create(c).Error
}

func Page(req *product.CategoryPageRes) (*product.CategoryRes, error) {}
