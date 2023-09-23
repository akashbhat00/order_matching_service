package models

import (
	"time"
)

type Products struct {
	ID        uint      `gorm:"column:id;primaryKey"`
	SellerID  uint      `gorm:"column:seller_id" json:"seller_id"`
	Name      string    `gorm:"column:name"`
	Price     float64   `gorm:"column:price"`
	Quantity  uint      `gorm:"column:quantity"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
