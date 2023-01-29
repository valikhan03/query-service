package kafka

import (
	"context"

	"search-service/repositories"

	"github.com/Shopify/sarama"
)

func NewConsumer() sarama.Consumer {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{}, config)
	if err != nil{
		
	}

	return consumer
}

type KafkaEventsConsumer struct{
	consumerGroup sarama.ConsumerGroup
}

func InitKafkaEventsConsumer(){

}

// must be executed as a goroutine
func (kc *KafkaEventsConsumer) Run() {
	handler := consumerGroupHandler{

	}
	
	go func(){
		for{
			select{
			case err := <- kc.consumerGroup.Errors():
				//log errors
			}
		}
	}()

	for{
		err := kc.consumerGroup.Consume(context.Background(), []string{}, &handler)
		if err != nil{

		}
	}
}

type consumerGroupHandler struct{
	ready chan bool
	repository *repositories.EventsRepository
}

func (c *consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	//mark the consumer as ready
	close(c.ready)
	return nil
}


func (c *consumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}


func (c *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for{
		select{
		case msg := <- claim.Messages():
			// read message
			// decode message
			// run needed function from repository
			session.MarkMessage(msg, "")
		case <- session.Context().Done():
			return nil
		}
	}
}