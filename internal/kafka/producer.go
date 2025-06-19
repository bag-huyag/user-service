package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

type UserEvent struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Type  string `json:"type"`
}

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokerAddr, topic string) *Producer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddr},
		Topic:   topic,
	})
	return &Producer{writer: writer}
}

func (p *Producer) SendUserEvent(event UserEvent) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = p.writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(event.ID),
		Value: data,
	})
	if err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
	}
	return err
}

func (p *Producer) Close() {
	p.writer.Close()
}
