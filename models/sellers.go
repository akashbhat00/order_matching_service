package models

import (
	"time"
)

type Sellers struct {
	ID        uint       `gorm:"column:id;primaryKey"`
	Name      string     `gorm:"column:name"`
	Email     string     `gorm:"column:email;unique"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	Products  []Products `gorm:"foreignKey:SellerID"`
}
