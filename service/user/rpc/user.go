package main

import (
	"flag"
	"fmt"
	"go-zero-micro/core/etcd"
	"go-zero-micro/core/kafka"
	"go-zero-micro/core/model"
	zeroConfig "go-zero-micro/service/user/rpc/internal/config"
	"go-zero-micro/service/user/rpc/internal/consumer"
	"go-zero-micro/service/user/rpc/internal/server"
	"go-zero-micro/service/user/rpc/internal/svc"
	"go-zero-micro/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c zeroConfig.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	etcd.Setup()
	model.Setup(etcd.C.Database)
	handler := &consumer.GroupHandler{}
	group := kafka.SetupConsumer(etcd.C.Kafka, handler)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer func() {
		s.Stop()
		_ = group.Close()
	}()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
