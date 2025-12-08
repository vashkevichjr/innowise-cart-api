package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/vashkevichjr/innowise-cart-api/internal/entity"
)

type CartRepo interface {
	CreateCart(ctx context.Context) (*entity.Cart, error)
	CreateItem(ctx context.Context, product string, price float32) (*entity.Item, error)
	AddItemToCart(ctx context.Context, cartId int32, itemId int32, quantity int32) error
	GetCart(ctx context.Context, id int32) (*entity.Cart, error)
	GetItem(ctx context.Context, id int32) (*entity.Item, error)
	GetItemsByCart(ctx context.Context, id int32) ([]entity.CartItem, error)
	SoftDeleteItemByCart(ctx context.Context, cartId int32, itemId int32) error
}

type Cart struct {
	repo CartRepo
}

func NewCart(repo CartRepo) *Cart {
	return &Cart{repo: repo}
}

// CONTRACT SERVICES

func (s *Cart) CreateCart(ctx context.Context) (cart *entity.Cart, err error) {
	row, err := s.repo.CreateCart(ctx)
	if err != nil {
		return nil, err
	}

	cart = &entity.Cart{
		Id:        row.Id,
		Items:     []entity.CartItem{},
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}

	return cart, nil
}

func (s *Cart) AddItemToCart(ctx context.Context, CartId int32, ItemId int32, Quantity int32) (cartItem *entity.CartItem, err error) {
	err = s.repo.AddItemToCart(ctx, CartId, ItemId, Quantity)
	if err != nil {
		return nil, err
	}

	item, err := s.repo.GetItem(ctx, ItemId)
	if err != nil {
		return nil, err
	}

	cartItem = &entity.CartItem{
		CartID:    CartId,
		ItemID:    ItemId,
		Name:      item.Name,
		Price:     item.Price,
		Quantity:  Quantity,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
	return cartItem, nil
}

func (s *Cart) CreateItem(ctx context.Context, product string, price float32) (item *entity.Item, err error) {
	row, err := s.repo.CreateItem(ctx, product, price)
	if err != nil {
		return nil, err
	}

	item = &entity.Item{
		ID:        row.ID,
		Name:      row.Name,
		Price:     row.Price,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}

	return item, nil
}

func (s *Cart) RemoveFromCart(ctx context.Context, CartId int32, ItemId int32) error {
	err := s.repo.SoftDeleteItemByCart(ctx, CartId, ItemId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Cart) ViewCart(ctx context.Context, id int32) (cart *entity.Cart, err error) {
	row, err := s.repo.GetCart(ctx, id)
	if err != nil {
		return nil, err
	}
	items, err := s.repo.GetItemsByCart(ctx, id)
	if err != nil {
		return nil, err
	}

	cart = &entity.Cart{
		Id:        id,
		Items:     items,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}

	return cart, nil
}

func (s *Cart) CalculatePrice(ctx context.Context, id int32) (calculator *entity.Calculator, err error) {
	cart, err := s.ViewCart(ctx, id)
	if err != nil {
		return nil, err
	}

	var totalPrice float32
	for _, item := range cart.Items {
		totalPrice += item.Price * float32(item.Quantity)
	}

	var discount int32 = 0

	if len(cart.Items) == 0 {
		return nil, errors.New("no items in cart")
	} else if totalPrice > 5000 {
		discount = 10
	} else if len(cart.Items) >= 3 {
		discount = 5
	} else {
		discount = 1
	}

	calculator = &entity.Calculator{
		CartID:          cart.Id,
		TotalPrice:      totalPrice,
		DiscountPercent: discount,
		FinalPrice:      totalPrice * (1 - float32(discount)/100),
	}

	return calculator, nil
}
