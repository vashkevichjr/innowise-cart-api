package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
	transport "github.com/vashkevichjr/innowise-cart-api/internal/transport/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vashkevichjr/innowise-cart-api/internal/config"
	"github.com/vashkevichjr/innowise-cart-api/internal/repository"
	"github.com/vashkevichjr/innowise-cart-api/internal/service"
	"github.com/vashkevichjr/innowise-cart-api/pkg/postgres"
)

type App struct {
	pool *pgxpool.Pool
	srv  *http.Server
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

	repoCart := repository.NewCartRepo(pool)

	serviceCart := service.NewCart(repoCart)

	handlerCart := transport.NewCartHandler(serviceCart)

	r := gin.Default()
	r.POST("/carts", handlerCart.CreateCart)
	r.POST("/items", handlerCart.CreateItem)
	r.POST("/carts/:cart_id/items/:item_id", handlerCart.AddItemToCart)
	r.DELETE("/carts/:cart_id/items/:item_id", handlerCart.RemoveItemFromCart)
	r.GET("/carts/:id", handlerCart.ViewCart)
	r.GET("/carts/:id/price", handlerCart.CalculateCart)

	server := &http.Server{
		Addr:    cfg.HTTPPort,
		Handler: r,
	}

	return &App{pool: pool, srv: server}, err
}

func (a *App) Run(ctx context.Context) error {
	log.Printf("Starting app on %s", a.srv.Addr)
	err := a.srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}

func (a *App) Close(ctx context.Context) error {
	log.Println("Shutting down...")
	if err := a.srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown error: %w", err)
	}

	if a.pool != nil {
		a.pool.Close()
	}
	return nil
}
