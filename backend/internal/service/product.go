package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/00Chaotic/flip-tech-test/backend/internal/model"
)

type ProductRepository interface {
	GetProducts(ctx context.Context) ([]*model.Product, error)
	GetProductBySKU(ctx context.Context, sku string) (*model.Product, error)
	UpdateProductInventory(ctx context.Context, sku string, difference int) error
}

type ProductService struct {
	productRepository ProductRepository
}

func NewProductService(productRepository ProductRepository) ProductService {
	return ProductService{productRepository: productRepository}
}

func (s *ProductService) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer ctx.Done()

	products, err := s.productRepository.GetProducts(ctx)
	if err != nil {
		log.Println("failed to get products", err)
		http.Error(w, "failed to get products", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(products)
	if err != nil {
		log.Println("failed to marshal json", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.Write(res)
}
