package product_model

import "go-zero-micro/core/model"

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
