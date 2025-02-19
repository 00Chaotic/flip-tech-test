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

func (s *ProductService) PurchaseProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer ctx.Done()

	var req model.PurchaseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("invalid request body", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	totalPrice := 0.0

	for _, item := range req.Items {
		product, err := s.productRepository.GetProductBySKU(ctx, item.SKU)
		if err != nil {
			log.Println("failed to get product by SKU", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		totalPrice += product.Price * float64(item.Quantity)

		err = s.productRepository.UpdateProductInventory(ctx, item.SKU, -item.Quantity)
		if err != nil {
			log.Println("failed to update product inventory", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}

	res := model.PurchaseResponse{TotalPrice: totalPrice}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println("failed to encode response", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}
