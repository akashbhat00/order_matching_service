package dto

type Product struct {
	ID       string `gorm:"column:product_id" json:"product_id"`
	Name     string `gorm:"column:product_name" json:"product_name"`
	Price    string `gorm:"column:product_price" json:"product_price"`
	Quantity string `gorm:"column:product_quantity" json:"product_quantity"`
}

type ProductRequestDto struct {
	Name     string  `gorm:"column:product_name" json:"product_name"`
	Price    float64 `gorm:"column:product_price" json:"product_price"`
	Quantity uint    `gorm:"column:product_quantity" json:"product_quantity"`
}
