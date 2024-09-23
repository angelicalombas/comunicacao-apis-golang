package models

import "time"

type Order struct {
	ID              uint       `json:"id" gorm:"primaryKey"`
	UserID          uint       `json:"user_id" validate:"required"`
	ItemDescription string     `json:"item_description" validate:"required"`
	ItemQuantity    int        `json:"item_quantity" validate:"required"`
	ItemPrice       float64    `json:"item_price" validate:"required"`
	TotalValue      float64    `json:"total_value" validate:"required"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty" gorm:"default:null"`
}

type OrderRequest struct {
	UserID          uint    `json:"user_id" validate:"required"`
	ItemDescription string  `json:"item_description" validate:"required"`
	ItemQuantity    int     `json:"item_quantity" validate:"required"`
	ItemPrice       float64 `json:"item_price" validate:"required"`
	TotalValue      float64 `json:"total_value" validate:"required"`
}
