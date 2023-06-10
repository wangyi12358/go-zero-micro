package logic

import (
	"context"
	"go-zero-micro/service/product/model/product_model"

	"go-zero-micro/service/product/rpc/internal/svc"
	"go-zero-micro/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProductLogic) CreateProduct(in *product.CreateProductReq) (*product.ProductRes, error) {
	p := &product_model.Product{
		Name:       in.Name,
		URL:        in.Url,
		CategoryID: in.CategoryId,
	}
	err := product_model.Create(p)
	if err != nil {
		return nil, err
	}
	return product_model.OfProductResponse(p), nil
}
