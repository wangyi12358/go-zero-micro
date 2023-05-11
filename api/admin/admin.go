package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-micro/common/errorx"
	"go-zero-micro/core/etcd"
	"go-zero-micro/core/kafka"
	"go-zero-micro/core/redis"
	"net/http"

	"go-zero-micro/api/admin/internal/config"
	"go-zero-micro/api/admin/internal/handler"
	"go-zero-micro/api/admin/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/admin.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	etcd.Setup()
	redis.Setup(etcd.C.CacheRedis)
	p := kafka.SetupProducer(etcd.C.Kafka)
	server := rest.MustNewServer(c.RestConf)
	defer func() {
		server.Stop()
		_ = p.Close()
	}()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	httpx.SetErrorHandlerCtx(func(ctx context.Context, err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
