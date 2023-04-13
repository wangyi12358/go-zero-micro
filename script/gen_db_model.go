package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/exec"
)

type GenConfig struct {
	Dsn    string `yaml:"dsn"`
	Db     string `yaml:"db"`
	Tables string `yaml:"tables"`
}

func getServiceOptions() []string {
	services, err := ioutil.ReadDir("./service")
	if err != nil {
		panic(fmt.Sprintf("Get services error: %s", err.Error()))
	}

	options := make([]string, len(services))
	for index, service := range services {
		options[index] = service.Name()
	}

	return options
}

func getGenConfig() GenConfig {
	data, err := ioutil.ReadFile("./gen.yaml")
	if err != nil {
		panic(fmt.Sprintf("Get gen config error: %s", err.Error()))
	}
	var config GenConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(fmt.Sprintf("Gen config conversion error: %s", err.Error()))
	}
	return config
}

func main() {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("err: %s\n", err)
		}
	}()

	genConfig := getGenConfig()

	cmd := exec.Command("gentool",
		"-dsn", genConfig.Dsn,
		"-tables", genConfig.Tables,
		"-db", genConfig.Db,
		"-outPath", "./pkg/models/gen")
	_, err := cmd.Output()

	if err != nil {
		panic(fmt.Sprintf("Command error: %s", err.Error()))
	}

	fmt.Printf("\033[32m%s\033[0m\n", "Success!")
}
