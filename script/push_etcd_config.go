package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"go-zero-micro/core/etcd"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type FlagConfig struct {
}

func getConfig(env string) (*etcd.Config, error) {
	var conf etcd.Config
	viper.SetEnvPrefix(env)
	viper.AutomaticEnv()
	viper.SetConfigName("config." + env)
	viper.AddConfigPath("./config/")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	if err = viper.Unmarshal(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}

func getFlagValue() (string, string) {
	// Get command line arguments
	env := flag.String("env", "dev", "environment variable")
	host := flag.String("host", "localhost", "etcd ip address")
	flag.Parse()
	return *env, *host
}

func main() {
	env, host := getFlagValue()
	// Get local config
	conf, err := getConfig(env)
	if err != nil {
		panic(fmt.Sprintf("get config error: %v", err))
	}
	fmt.Println(conf)
	// Connect to Etcd cluster
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{fmt.Sprintf("http://%s:2379", host)},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// Put a key-value pair into Etcd
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	value, err := json.Marshal(conf)
	if err != nil {
		panic(fmt.Sprintf("Json marshal error: %v", err))
	}
	_, err = cli.Put(ctx, etcd.ConfigKey, string(value))
	if err != nil {
		panic(fmt.Sprintf("Etcd configuration push failed: %v", err))
	}
	fmt.Printf("\033[32m%s\033[0m\n", "Succeeded in pushing Etcd configuration")
}
