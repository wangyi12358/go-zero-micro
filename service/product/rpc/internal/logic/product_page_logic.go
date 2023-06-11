package logic

import (
	"context"
	"go-zero-micro/service/product/model/product_model"

	"go-zero-micro/service/product/rpc/internal/svc"
	"go-zero-micro/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductPageLogic {
	return &ProductPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductPageLogic) ProductPage(in *product.ProductPageReq) (*product.ProductPageRes, error) {
	res, err := product_model.Page(in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
