package entity

import "time"

type Cart struct {
	Id        int32
	Items     []CartItem
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
