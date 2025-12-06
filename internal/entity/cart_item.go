package entity

import "time"

type CartItem struct {
	CartID    int32      `json:"cart_id"`
	ItemID    int32      `json:"item_id"`
	Quantity  int32      `json:"quantity"`
	Name      string     `json:"name"`
	Price     float32    `json:"price"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
