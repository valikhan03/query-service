package kafka

import(
	"github.com/Shopify/sarama"
)

func NewConsumer() sarama.ConsumerGroup {
	config := sarama.NewConfig()
	consumerGroup, err := sarama.NewConsumerGroup([]string{}, "", config)
	if err != nil{

	}

	return consumerGroup
}