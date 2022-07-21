package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/awesome-sphere/as-seating-consumer/kafka/interfaces"
	"github.com/awesome-sphere/as-seating-consumer/redis"
	"github.com/segmentio/kafka-go"
)

func initReader(topic_name string, groupBalancers []kafka.GroupBalancer) *kafka.Reader {
	config := kafka.ReaderConfig{
		Brokers:        []string{KAFKA_LOCATION},
		Topic:          topic_name,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
		GroupID:        "seating-status-consumer",
		GroupBalancers: groupBalancers,
	}
	reader := kafka.NewReader(config)
	return reader
}

func readFromReader(reader *kafka.Reader) {
	defer reader.Close()
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Error while reading message: %v", err)
			continue
		}
		var value interfaces.WriterInterface
		err = json.Unmarshal(msg.Value, &value)

		if err != nil {
			log.Fatalf("Failed to unmarshal message: %v", err.Error())
			continue
		}

		redis.UpdateStatus(value.TheaterID, value.TimeSlotID, value.SeatID, value.SeatStatus)
	}
}

func ReadMessage(topic_name string) {
	groupBalancers := make([]kafka.GroupBalancer, 0)
	groupBalancers = append(groupBalancers, kafka.RangeGroupBalancer{})

	readers := make([]*kafka.Reader, 0)
	for i := 0; i < PARTITION; i++ {
		readers = append(readers, initReader(topic_name, groupBalancers))
	}
	for _, reader := range readers {
		go readFromReader(reader)
	}
}
