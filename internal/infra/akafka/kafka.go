package akafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

func Consume(topics []string, servers string, msgChan chan *kafka.Message) {
	kafkaCosumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
		"groud.id":          "go-api",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	kafkaCosumer.SubscribeTopics(topics, nil)

	for {
		msg, err := kafkaCosumer.ReadMessage(-1)
		if err != nil {
			msgChan <- msg
		}
	}
}
