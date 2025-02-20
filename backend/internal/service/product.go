package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/00Chaotic/flip-tech-test/backend/internal/model"
)

type ProductRepository interface {
	GetProducts(ctx context.Context) ([]*model.Product, error)
	UpdateProductInventories(ctx context.Context, items []model.PurchaseItem) (float64, []*model.Product, error)
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

	res := model.ProductsResponse{Products: products}
	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.Println("failed to encode response", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
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

	for _, item := range req.Items {
		if item.Quantity < 0 {
			http.Error(w, fmt.Sprintf("item quantity is negative for SKU: %s", item.SKU), http.StatusBadRequest)
			return
		}
	}

	totalPrice, updatedProducts, err := s.productRepository.UpdateProductInventories(ctx, req.Items)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "one or more items unavailable", http.StatusConflict)
			return
		}

		log.Println("failed to update product inventories", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	res := model.PurchaseResponse{TotalPrice: totalPrice, UpdatedProducts: updatedProducts}
	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.Println("failed to encode response", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}
