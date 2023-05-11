package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"go-zero-micro/core/kafka"
	"go-zero-micro/core/model"
	clientv3 "go.etcd.io/etcd/client/v3"
	"os"
	"time"
)

type Config struct {
	Database   model.Database
	Kafka      kafka.Config
	CacheRedis cache.CacheConf
}

const ConfigKey = "config"

var C Config

func Setup() {
	addr := os.Getenv("ETCD_ADDR")
	if addr == "" {
		addr = "localhost:2379"
	}
	// Connect to Etcd cluster
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("http://%s", addr)},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	// Put a key-value pair into Etcd
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := cli.Get(ctx, ConfigKey)
	if err != nil {
		fmt.Printf("Etcd get config error: %v\n", err)
	}
	err = json.Unmarshal(res.Kvs[0].Value, &C)
	if err != nil {
		fmt.Printf("Etcd config unmarshal error: %v\n", err)
	}
}
