package main

import (
	"context"
	"log"

	"github.com/00Chaotic/flip-tech-test/backend/internal/config"
	"github.com/00Chaotic/flip-tech-test/backend/internal/wiring"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	wiring.StartServer(ctx, cfg)
}
