package models

import (
	"time"
)

type Orders struct {
	ID         uint      `gorm:"column:id;primaryKey"`
	BuyerID    uint      `gorm:"column:buyer_id"`
	ProductID  uint      `gorm:"column:product_id"`
	Quantity   uint      `gorm:"column:quantity"`
	TotalPrice float64   `gorm:"column:total_price"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}
