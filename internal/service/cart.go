package service

import (
	"context"

	"github.com/vashkevichjr/innowise-cart-api/internal/repository"
)

type CartService struct {
	repo repository.Cart
}

// изменить нейминг на Cart
func NewCartService(repo repository.Cart) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) CreateCart(ctx context.Context) {
	s.repo.CreateCart(ctx)
}
