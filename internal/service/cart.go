package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/vashkevichjr/innowise-cart-api/internal/db"
	"github.com/vashkevichjr/innowise-cart-api/internal/entity"
	"github.com/vashkevichjr/innowise-cart-api/internal/repository"
)

type Cart struct {
	repo *repository.Cart
}

func NewCart(repo *repository.Cart) *Cart {
	return &Cart{repo: repo}
}

// CONTRACT SERVICES

func (s *Cart) CreateCart(ctx context.Context) (cart *entity.Cart, err error) {
	row, err := s.repo.CreateCart(ctx)
	if err != nil {
		return nil, err
	}

	cart = &entity.Cart{
		Id:        row.ID,
		Items:     []entity.CartItem{},
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}

	return cart, nil
}

func (s *Cart) AddItemToCart(ctx context.Context, CartId int32, ItemId int32, Quantity int32) (cartItem *entity.CartItem, err error) {
	args := db.AddItemToCartParams{
		CartID:   CartId,
		ItemID:   ItemId,
		Quantity: Quantity,
	}

	err = s.repo.AddItemToCart(ctx, args)
	if err != nil {
		return nil, err
	}

	item, err := s.repo.GetItem(ctx, args.ItemID)
	if err != nil {
		return nil, err
	}

	cartItem = &entity.CartItem{
		CartID:    args.CartID,
		ItemID:    args.ItemID,
		Name:      item.Product,
		Price:     item.Price,
		Quantity:  args.Quantity,
		CreatedAt: item.CreatedAt.Time,
		UpdatedAt: item.UpdatedAt.Time,
	}
	return cartItem, nil
}

func (s *Cart) CreateItem(ctx context.Context, product string, price float32) (item *entity.Item, err error) {
	args := db.CreateItemParams{
		Product: product,
		Price:   price,
	}
	row, err := s.repo.CreateItem(ctx, args)
	if err != nil {
		return nil, err
	}

	item = &entity.Item{
		ID:        row.ID,
		Name:      row.Product,
		Price:     row.Price,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}

	return item, nil
}

func (s *Cart) RemoveFromCart(ctx context.Context, CartId int32, ItemId int32) error {
	args := db.SoftDeleteItemByCartParams{
		CartID: CartId,
		ItemID: ItemId,
	}

	err := s.repo.SoftDeleteItemByCart(ctx, args)
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
	rows, err := s.repo.GetItemsByCart(ctx, id)
	if err != nil {
		return nil, err
	}

	var items []entity.CartItem

	for _, i := range rows {
		items = append(items, entity.CartItem{CartID: i.CartID, ItemID: i.ItemID, Price: i.Price, Quantity: i.Quantity})
	}

	cart = &entity.Cart{
		Id:        id,
		Items:     items,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
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
