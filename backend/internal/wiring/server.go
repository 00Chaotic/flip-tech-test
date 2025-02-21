package wiring

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/00Chaotic/flip-tech-test/backend/internal/config"
	"github.com/00Chaotic/flip-tech-test/backend/internal/postgres"
	"github.com/00Chaotic/flip-tech-test/backend/internal/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func StartServer(ctx context.Context, cfg *config.Config) {
	dbx, err := sqlx.ConnectContext(ctx, "postgres", cfg.DatabaseUrl)
	if err != nil {
		log.Fatal("failed connecting to database", err)
	}
	defer dbx.Close()

	productDAO := postgres.NewProductDAO(dbx)

	productService := service.NewProductService(productDAO)

	router := mux.NewRouter()
	router.HandleFunc("/products", productService.GetProducts)
	router.HandleFunc("/purchase", productService.PurchaseProducts).Methods(http.MethodPut)

	// Only needed if running frontend outside Docker
	c := cors.New(cors.Options{
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedOrigins: []string{"http://localhost:*"},
	})

	handler := c.Handler(router)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		fmt.Println("Serving HTTP server")

		if err = srv.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			log.Fatalf("server listener error: %s\n", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		log.Println("Context cancelled, shutting down server...")
	case <-stop:
		log.Println("Interrupt signal received, shutting down server...")
	}

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server shutdown failed: %s\n", err)
	}

	log.Println("Server has been shut down")
}
