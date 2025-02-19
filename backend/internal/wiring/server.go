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

	log.Fatal(http.ListenAndServe(":8080", mux))
}
