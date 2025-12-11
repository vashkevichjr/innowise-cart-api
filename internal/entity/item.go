package entity

import "time"

type Item struct {
	ID        int32      `json:"id"`
	Name      string     `json:"name"`
	Price     float32    `json:"price"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
