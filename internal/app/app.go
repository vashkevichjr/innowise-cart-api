package app

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vashkevichjr/innowise-cart-api/internal/config"
	"github.com/vashkevichjr/innowise-cart-api/internal/repository"
	"github.com/vashkevichjr/innowise-cart-api/pkg/postgres"
)

type App struct {
	pool *pgxpool.Pool
	//srv  *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("error load config: %w", err)
	}

	pool, err := postgres.NewClient(ctx, cfg.PGDSN)
	if err != nil {
		return nil, fmt.Errorf("error create postgres client: %w", err)
	}

	repo := repository.NewCartRepo(pool)
	_ = repo

	return &App{pool: pool}, err
}

func (a *App) Run(ctx context.Context) error {
	log.Println("Starting app...")
	return nil
}

func (a *App) Close(ctx context.Context) error {
	log.Println("Shutting down...")
	if a.pool != nil {
		a.pool.Close()
	}
	return nil
}
