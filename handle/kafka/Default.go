package kafka

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaConfig struct {
	Url   string
	Topic string
}

type PhoneConfig struct {
	PhoneNumber string
}

func GetDefaultPhone(c *fiber.Ctx) (err error) {
	id := c.Params("id")

	cfgReq := KafkaConfig{
		Url:   "localhost:9092",
		Topic: "shopReq",
	}

	cfgRes := KafkaConfig{
		Url:   "localhost:9092",
		Topic: "shopRes",
	}

	connReq, err := kafka.DialLeader(context.Background(), "tcp", cfgReq.Url, cfgReq.Topic, 0)
	if err != nil {
		log.Fatal(err)
	}
	//

	connRes, err := kafka.DialLeader(context.Background(), "tcp", cfgRes.Url, cfgRes.Topic, 0)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := connReq.Close(); err != nil {
			log.Fatal(err)
		}
		if err := connRes.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	data := func() []kafka.Message {
		message := make([]kafka.Message, 0)
		message = append(message, kafka.Message{Value: []byte(id)})

		return message
	}()

	_, err = connReq.WriteMessages(data...)

	if err != nil {
		log.Fatal(err)
	}

	if _, err := connRes.Seek(0, kafka.SeekEnd); err != nil {
		log.Fatal(err)
	}

	for {
		message, err := connRes.ReadMessage(10e3)
		if err != nil {
			panic(err)
		}
		content := Response{
			Error: "",
			Name:  string(message.Value),
		}

		return c.JSON(content)
	}
}
