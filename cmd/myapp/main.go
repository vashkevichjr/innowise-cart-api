package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vashkevichjr/innowise-cart-api/internal/config"
	"github.com/vashkevichjr/innowise-cart-api/internal/repository"
	"github.com/vashkevichjr/innowise-cart-api/pkg/postgres"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	pool, err := postgres.NewClient(ctx, cfg.PGDSN)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	cartRepo := repository.NewCartRepo(pool)
	_ = cartRepo
	fmt.Println("starting server")
}
