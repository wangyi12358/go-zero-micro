package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &product.ProductPageRes{}, nil
}
