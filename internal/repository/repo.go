package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vashkevichjr/innowise-cart-api/internal/db"
)

type CartRepo struct {
	queries *db.Queries
	pool    *pgxpool.Pool
}

func NewCartRepo(pool *pgxpool.Pool) *CartRepo {
	return &CartRepo{
		queries: db.New(pool),
		pool:    pool,
	}
}
