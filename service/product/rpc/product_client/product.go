// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package product_client

import (
	"context"

	"go-zero-micro/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CategoryPageRes  = product.CategoryPageRes
	CategoryReq      = product.CategoryReq
	CategoryRes      = product.CategoryRes
	CreateProductReq = product.CreateProductReq
	ProductPageReq   = product.ProductPageReq
	ProductPageRes   = product.ProductPageRes
	ProductRes       = product.ProductRes

	Product interface {
		CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*ProductRes, error)
		ProductPage(ctx context.Context, in *ProductPageReq, opts ...grpc.CallOption) (*ProductPageRes, error)
		CreateCategory(ctx context.Context, in *CategoryReq, opts ...grpc.CallOption) (*CategoryRes, error)
		CategoryPage(ctx context.Context, in *ProductPageReq, opts ...grpc.CallOption) (*ProductPageRes, error)
	}

	defaultProduct struct {
		cli zrpc.Client
	}
)

func NewProduct(cli zrpc.Client) Product {
	return &defaultProduct{
		cli: cli,
	}
}

func (m *defaultProduct) CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*ProductRes, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.CreateProduct(ctx, in, opts...)
}

func (m *defaultProduct) ProductPage(ctx context.Context, in *ProductPageReq, opts ...grpc.CallOption) (*ProductPageRes, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.ProductPage(ctx, in, opts...)
}

func (m *defaultProduct) CreateCategory(ctx context.Context, in *CategoryReq, opts ...grpc.CallOption) (*CategoryRes, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.CreateCategory(ctx, in, opts...)
}

func (m *defaultProduct) CategoryPage(ctx context.Context, in *ProductPageReq, opts ...grpc.CallOption) (*ProductPageRes, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.CategoryPage(ctx, in, opts...)
}
