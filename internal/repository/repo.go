package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vashkevichjr/innowise-cart-api/internal/db"
)

type Cart struct {
	*db.Queries
	pool *pgxpool.Pool
}

func NewCartRepo(pool *pgxpool.Pool) *Cart {
	return &Cart{
		Queries: db.New(pool),
		pool:    pool,
	}
}
