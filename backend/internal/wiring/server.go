package wiring

import (
	"context"
	"log"
	"net/http"

	"github.com/00Chaotic/flip-tech-test/backend/internal/config"
	"github.com/00Chaotic/flip-tech-test/backend/internal/postgres"
	"github.com/00Chaotic/flip-tech-test/backend/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func StartServer(ctx context.Context, cfg *config.Config) {
	dbx, err := sqlx.ConnectContext(ctx, "postgres", cfg.DatabaseUrl)
	if err != nil {
		log.Fatal("failed connecting to database", err)
	}

	productDAO := postgres.NewProductDAO(dbx)

	productService := service.NewProductService(productDAO)

	mux := http.NewServeMux()
	mux.HandleFunc("/products", productService.GetProducts)
	mux.HandleFunc("/purchase", productService.PurchaseProducts)

	// Only needed if running frontend outside Docker
	c := cors.New(cors.Options{
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedOrigins: []string{"http://localhost:*"},
	})

	handler := c.Handler(mux)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
