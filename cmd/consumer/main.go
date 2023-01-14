package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

// RecordValue represents the struct of the value in a Kafka message

func main() {
	consumer := NewKafkaConsumer()
	fmt.Println("start consuming")
	defer consumer.Close()
	var ctx = context.Background()

	Consume(consumer, ctx)
}

func NewKafkaConsumer() *kafka.Reader {
	return kafka.NewReader(
		kafka.ReaderConfig{
			Brokers: []string{"localhost:29092"},
			//Brokers:     []string{"localhost:29092"},
			Topic:       "simple",
			GroupID:     "kafka-group-id",
			StartOffset: kafka.LastOffset,
		},
	)
}

func Consume(consumer *kafka.Reader, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context canceled, stopping consumer")
			return
		default:
			msg, err := consumer.ReadMessage(ctx)
			if err != nil {
				fmt.Println("Error when consuming message: ", err)
				continue
			}
			fmt.Println("Recieved message: ", string(msg.Value))
		}
	}
}
