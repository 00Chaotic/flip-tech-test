package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	http2 "github.com/00Chaotic/flip-tech-test/backend/internal/http"
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
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	var res model.ProductsResponse

	products, err := s.productRepository.GetProducts(ctx)
	if err != nil {
		log.Println("failed to get products", err)

		res.Error = "failed to get products"
		http2.SendJSONResponse(w, res, http.StatusInternalServerError)
		return
	}

	res.Products = products
	http2.SendJSONResponse(w, res, http.StatusOK)
}

func (s *ProductService) PurchaseProducts(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	var req model.PurchaseRequest
	var res model.PurchaseResponse

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("invalid request body", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	for _, item := range req.Items {
		if item.Quantity < 0 {
			res.Error = fmt.Sprintf("item quantity is negative for SKU: %s", item.SKU)
			http2.SendJSONResponse(w, res, http.StatusBadRequest)

			return
		}
	}

	totalPrice, updatedProducts, err := s.productRepository.UpdateProductInventories(ctx, req.Items)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			res.Error = "one or more items unavailable"
			http2.SendJSONResponse(w, res, http.StatusConflict)

			return
		}

		log.Println("failed to update product inventories", err)

		res.Error = "purchase failed"
		http2.SendJSONResponse(w, res, http.StatusInternalServerError)

		return
	}

	res.TotalPrice = totalPrice
	res.UpdatedProducts = updatedProducts

	http2.SendJSONResponse(w, res, http.StatusOK)
}
