package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-micro/core/model"
)

type Config struct {
	zrpc.RpcServerConf
	Database model.Database
}
