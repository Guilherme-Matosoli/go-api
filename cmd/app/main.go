package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Guilherme-Matosoli/go-api/internal/infra/akafka"
	"github.com/Guilherme-Matosoli/go-api/internal/infra/repository"
	"github.com/Guilherme-Matosoli/go-api/internal/infra/web"
	"github.com/Guilherme-Matosoli/go-api/internal/usecase"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306/products)")

	if err != nil {
		panic(err)
	}

	defer db.Close()
	repository := repository.NewProductRepositoryPg(db)
	createProductUsecase := usecase.NewCreateProductUseCase(repository)
	listProductUsecase := usecase.NewListProductsUseCase(repository)

	productHandlers := web.NewProductHandlers(createProductUsecase, listProductUsecase)

	r := chi.NewRouter()

	r.Post("/products", productHandlers.CreateProductHandler)
	r.Get("/products", productHandlers.ListProductsHandler)

	go http.ListenAndServe(":4000", r)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"products"}, "host.docker.internal:9094", msgChan)

	for msg := range msgChan {
		dto := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dto)

		if err != nil {
			println(err)
		}

		_, err = createProductUsecase.Execute(dto)

		if err != nil {

		}
	}
}
