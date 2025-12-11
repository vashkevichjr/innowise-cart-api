package entity

import "time"

type CartItem struct {
	CartID    int32
	ItemID    int32
	Quantity  int32
	Name      string
	Price     float32
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
