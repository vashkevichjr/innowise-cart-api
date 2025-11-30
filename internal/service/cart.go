package service

import (
	"context"

	"github.com/vashkevichjr/innowise-cart-api/internal/repository"
)

type CartService struct {
	repo repository.CartRepo
}

func NewCartService(repo repository.CartRepo) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) CreateCart(ctx context.Context) {
	_ = 1
}
