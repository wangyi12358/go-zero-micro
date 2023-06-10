package product_model

import (
	"go-zero-micro/core/model"
	"go-zero-micro/service/product/rpc/types/product"
)

func OfProductResponse(p *Product) *product.ProductRes {
	if p == nil {
		return nil
	}
	return &product.ProductRes{
		Id:         p.ID,
		Name:       p.Name,
		Url:        p.URL,
		CategoryId: p.CategoryID,
	}
}

func Create(product *Product) error {
	return model.DB.Create(product).Error
}

func FindOneById(id int64) (*Product, error) {
	var product = &Product{}
	err := model.DB.First(product, id).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func Page() {

}
