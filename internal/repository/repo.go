package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vashkevichjr/innowise-cart-api/internal/db"
	"github.com/vashkevichjr/innowise-cart-api/internal/entity"
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

func (c *Cart) CreateCart(ctx context.Context) (*entity.Cart, error) {
	row, err := c.Queries.CreateCart(ctx)
	if err != nil {
		return nil, err
	}
	cart := &entity.Cart{
		Id:        row.ID,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
		DeletedAt: &row.DeletedAt.Time,
	}
	return cart, nil
}

func (c *Cart) CreateItem(ctx context.Context, product string, price float32) (*entity.Item, error) {
	row, err := c.Queries.CreateItem(ctx, db.CreateItemParams{Product: product, Price: price})
	if err != nil {
		return nil, err
	}
	item := &entity.Item{
		ID:        row.ID,
		Name:      row.Product,
		Price:     row.Price,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
		DeletedAt: &row.DeletedAt.Time,
	}
	return item, nil
}

func (c *Cart) AddItemToCart(ctx context.Context, cartId int32, itemId int32, quantity int32) error {
	err := c.Queries.AddItemToCart(ctx, db.AddItemToCartParams{CartID: cartId, ItemID: itemId, Quantity: quantity})
	if err != nil {
		return err
	}
	return nil
}

func (c *Cart) GetCart(ctx context.Context, id int32) (*entity.Cart, error) {
	row, err := c.Queries.GetCart(ctx, id)
	if err != nil {
		return nil, err
	}
	cart := &entity.Cart{
		Id:        row.ID,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
		DeletedAt: &row.DeletedAt.Time,
	}
	return cart, nil
}

func (c *Cart) GetItem(ctx context.Context, id int32) (*entity.Item, error) {
	row, err := c.Queries.GetItem(ctx, id)
	if err != nil {
		return nil, err
	}
	item := &entity.Item{
		ID:        row.ID,
		Name:      row.Product,
		Price:     row.Price,
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}
	return item, nil
}

func (c *Cart) GetItemsByCart(ctx context.Context, id int32) ([]entity.CartItem, error) {
	rows, err := c.Queries.GetItemsByCart(ctx, id)
	if err != nil {
		return nil, err
	}

	items := make([]entity.CartItem, len(rows))

	for _, row := range rows {
		item := entity.CartItem{
			CartID:   row.CartID,
			ItemID:   row.ItemID,
			Quantity: row.Quantity,
			Name:     row.Product,
			Price:    row.Price,
		}
		items = append(items, item)
	}

	return items, nil
}

func (c *Cart) SoftDeleteItemByCart(ctx context.Context, cartId int32, itemId int32) error {
	err := c.Queries.SoftDeleteItemByCart(ctx, db.SoftDeleteItemByCartParams{CartID: cartId, ItemID: itemId})
	if err != nil {
		return err
	}
	return nil
}
