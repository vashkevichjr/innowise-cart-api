package service

import (
	"context"

	"github.com/vashkevichjr/innowise-cart-api/internal/db"
	"github.com/vashkevichjr/innowise-cart-api/internal/repository"
)

type CartService struct {
	repo *repository.Cart
}

func NewCart(repo *repository.Cart) *CartService {
	return &CartService{repo: repo}
}

//CARTS SERVICES

func (s *CartService) CreateCart(ctx context.Context) (int32, error) {
	id, err := s.repo.CreateCart(ctx)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *CartService) GetCart(ctx context.Context, id int32) (*db.Cart, error) {
	cart, err := s.repo.GetCart(ctx, id)
	if err != nil {
		return nil, err
	}
	return &cart, nil
}

func (s *CartService) GetCarts(ctx context.Context) (*[]db.Cart, error) {
	carts, err := s.repo.GetCarts(ctx)
	if err != nil {
		return nil, err
	}
	return &carts, nil
}

func (s *CartService) SoftDeleteCart(ctx context.Context, id int32) error {
	err := s.repo.SoftDeleteCart(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *CartService) HardDeleteCart(ctx context.Context, id int32) error {
	err := s.repo.HardDeleteCart(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

//ITEMS SERVICES

func (s *CartService) CreateItem(ctx context.Context, params db.CreateItemParams) (int32, error) {
	id, err := s.repo.CreateItem(ctx, params)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *CartService) GetItem(ctx context.Context, id int32) (*db.GetItemRow, error) {
	item, err := s.repo.GetItem(ctx, id)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *CartService) GetItems(ctx context.Context) (*[]db.GetItemsRow, error) {
	items, err := s.repo.GetItems(ctx)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func (s *CartService) UpdateItem(ctx context.Context, params db.UpdateItemParams) (int32, error) {
	id, err := s.repo.UpdateItem(ctx, params)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *CartService) UpdateItemPrice(ctx context.Context, params db.UpdateItemPriceParams) (int32, error) {
	id, err := s.repo.UpdateItemPrice(ctx, params)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *CartService) UpdateItemProduct(ctx context.Context, params db.UpdateItemProductParams) (int32, error) {
	id, err := s.repo.UpdateItemProduct(ctx, params)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *CartService) SoftDeleteItem(ctx context.Context, id int32) error {
	err := s.repo.SoftDeleteItem(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *CartService) HardDeleteItem(ctx context.Context, id int32) error {
	err := s.repo.HardDeleteItem(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

//CART-ITEMS

func (s *CartService) AddItemToCart(ctx context.Context, params db.AddItemToCartParams) error {
	err := s.repo.AddItemToCart(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (s *CartService) UpdateItemInCart(ctx context.Context, params db.UpdateItemInCartParams) error {
	err := s.repo.UpdateItemInCart(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (s *CartService) UpdateItemInCartQuantity(ctx context.Context, params db.UpdateItemInCartQuantityParams) error {
	err := s.repo.UpdateItemInCartQuantity(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (s *CartService) GetItemsByCartID(ctx context.Context, cartID int32) (*[]db.GetItemsByCartRow, error) {
	items, err := s.repo.GetItemsByCart(ctx, cartID)
	if err != nil {
		return nil, err
	}
	return &items, err
}

func (s *CartService) GetCartsByItem(ctx context.Context, cartID int32) (*[]db.GetCartsByItemRow, error) {
	items, err := s.repo.GetCartsByItem(ctx, cartID)
	if err != nil {
		return nil, err
	}
	return &items, err
}

func (s *CartService) GetCartsItems(ctx context.Context) (*[]db.GetCartsItemsRow, error) {
	items, err := s.repo.GetCartsItems(ctx)
	if err != nil {
		return nil, err
	}
	return &items, err
}

func (s *CartService) SoftDeleteItemByCart(ctx context.Context, params db.SoftDeleteItemByCartParams) error {
	err := s.repo.SoftDeleteItemByCart(ctx, params)
	if err != nil {
		return err
	}
	return nil
}

func (s *CartService) HardDeleteItemByCart(ctx context.Context, params db.HardDeleteItemByCartParams) error {
	err := s.repo.HardDeleteItemByCart(ctx, params)
	if err != nil {
		return err
	}
	return nil
}
