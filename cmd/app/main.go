package main

import (
	"database/sql"

	"github.com/Guilherme-Matosoli/go-api/internal/infra/akafka"
	"github.com/Guilherme-Matosoli/go-api/internal/infra/repository"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306/products)")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	msgChan := make(chan *kafka.Message)
	akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	repository := repository.NewProductRepositoryPg(db)
}
