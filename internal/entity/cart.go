package entity

import "time"

type Cart struct {
	Id        int32      `json:"id"`
	Items     []CartItem `json:"items,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
