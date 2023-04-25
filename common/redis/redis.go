package redis

import (
	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/cache"
)

var DB *goredis.Client

func Setup(c cache.CacheConf) {
	DB = goredis.NewClient(&goredis.Options{
		Addr:     c[0].Host,
		Password: c[0].Pass,
	})
}
