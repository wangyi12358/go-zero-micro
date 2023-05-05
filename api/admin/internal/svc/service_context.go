package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-micro/api/admin/internal/config"
	"go-zero-micro/service/product/rpc/product_client"
	"go-zero-micro/service/user/rpc/user_client"
)

type ServiceContext struct {
	Config     config.Config
	UserRpc    user_client.User
	ProductRpc product_client.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserRpc:    user_client.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc: product_client.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
