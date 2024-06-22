package web

import (
	"encoding/json"
	"net/http"

	"github.com/Guilherme-Matosoli/go-api/internal/usecase"
)

type ProductHandlers struct {
	CreateProductsUseCase *usecase.CreateProductUseCase
	ListProductsUseCase   *usecase.ListProductsUseCase
}

func NewProductHandlers(createProductUsecase *usecase.CreateProductUseCase, listProductsUsecase *usecase.ListProductsUseCase) *ProductHandlers {
	return &ProductHandlers{
		createProductUsecase,
		listProductsUsecase,
	}
}

func (p *ProductHandlers) CreateProductHandler(w http.ResponseWriter, r http.Request) {
	var input usecase.CreateProductInputDto
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	output, err := p.CreateProductsUseCase.Execute(input)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)

}

func (p *ProductHandlers) ListProductsHandler(w http.ResponseWriter, r http.Request) {
	output, err := p.ListProductsUseCase.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}
