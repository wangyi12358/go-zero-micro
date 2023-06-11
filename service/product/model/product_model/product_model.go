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

func OfProductResponses(products []Product) []*product.ProductRes {
	var list = make([]*product.ProductRes, len(products))
	for i, u := range products {
		list[i] = OfProductResponse(&u)
	}
	return list
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

func Page(req *product.ProductPageReq) (res *product.ProductPageRes, err error) {
	var products []Product
	var total int64
	offset := req.PageSize * (req.Current - 1)
	err = model.DB.Model(&Product{}).Offset(int(offset)).Limit(int(req.PageSize)).Order("id desc").Find(&products).Error
	if err != nil {
		return
	}
	err = model.DB.Model(&Product{}).Count(&total).Error
	if err != nil {
		return
	}
	return &product.ProductPageRes{
		Total: int32(total),
		List:  OfProductResponses(products),
	}, nil
}
