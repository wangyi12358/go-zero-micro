package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"log"
)

type Config struct {
	Brokers  []string
	ClientId string
	Consumer struct {
		Topics  []string
		GroupId string
	}
}

var Producer sarama.SyncProducer
var Consumer sarama.ConsumerGroup

func SetupProducer(c Config) sarama.SyncProducer {
	var err error
	// 配置Kafka生产者
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	// 连接Kafka
	Producer, err = sarama.NewSyncProducer(c.Brokers, config)
	if err != nil {
		log.Fatalln("Failed to connect Kafka:", err)
	}

	return Producer
}

func SetupConsumer(c Config, handler sarama.ConsumerGroupHandler) sarama.ConsumerGroup {
	var err error
	// 配置Kafka消费者
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// 连接Kafka
	Consumer, err = sarama.NewConsumerGroup(c.Brokers, c.Consumer.GroupId, config)
	if err != nil {
		log.Fatalln("Failed to connect Kafka:", err)
	}
	//defer consumer.Close()
	ctx := context.Background()
	go func() {
		for {
			err = Consumer.Consume(ctx, c.Consumer.Topics, handler)
			if err != nil {
				log.Printf("Error from consumer: %v", err)
			}
		}
	}()

	return Consumer
}
