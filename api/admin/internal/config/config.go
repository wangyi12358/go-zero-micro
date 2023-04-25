package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-micro/common/kafka"
)

type Config struct {
	rest.RestConf
	UserRpc    zrpc.RpcClientConf
	CacheRedis cache.CacheConf
	Auth       struct {
		AccessSecret string
		AccessExpire int64
	}
	Kafka kafka.Config
}
