package es

import (
	"github.com/olivere/elastic"
	"log"
)

type Config struct {
	Url      string
	Username string
	Password string
}

var Client *elastic.Client

func Setup(c Config) {
	var err error
	// 创建ES client用于后续操作ES
	Client, err = elastic.NewClient(elastic.SetURL(c.Url), elastic.SetBasicAuth(c.Username, c.Password))
	if err != nil {
		log.Fatalf("es.Setup err: %v", err)
	}
}
