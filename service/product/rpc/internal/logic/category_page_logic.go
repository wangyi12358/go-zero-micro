package logic

import (
	"context"

	"go-zero-micro/service/product/rpc/internal/svc"
	"go-zero-micro/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryPageLogic {
	return &CategoryPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CategoryPageLogic) CategoryPage(in *product.ProductPageReq) (*product.ProductPageRes, error) {
	// todo: add your logic here and delete this line

	return &product.ProductPageRes{}, nil
}
