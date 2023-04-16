package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
	Port     string
}

var DatabaseConfig = &Database{}

func Setup() {
	wd, _ := os.Getwd()
	configDir := filepath.Join(wd, "configs/")
	fmt.Printf("configDir: %s", configDir)
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.AddConfigPath(configDir)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()

	if viper.UnmarshalKey("database", DatabaseConfig) != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
