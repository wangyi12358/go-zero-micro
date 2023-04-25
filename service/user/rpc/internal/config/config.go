package config

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-micro/common/kafka"
	"go-zero-micro/common/model"
)

type Config struct {
	zrpc.RpcServerConf
	Database model.Database
	Kafka    kafka.Config
}
