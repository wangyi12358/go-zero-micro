package consumer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"go-zero-micro/common/kafka"
)

type GroupHandler struct {
}

func (GroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	fmt.Println("Consumer setup")
	return nil
}

func (GroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	fmt.Println("Consumer cleanup")
	return nil
}

func (GroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	fmt.Printf("Consumer claim\n")

	for message := range claim.Messages() {
		value := string(message.Value)
		fmt.Printf("Consumer value: %s \n", value)
		if message.Topic == kafka.TestTopic {
			fmt.Println("test topic message")
		}

		session.MarkMessage(message, "")
	}

	return nil
}
