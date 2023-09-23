package dto

// CreateSellerRequest represents the request data for creating a seller.
type CreateSellerRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type SellerWithProducts struct {
	Name     string    `gorm:"column:name" json:"name"`
	Email    string    `gorm:"column:email" json:"email"`
	Products []Product `gorm:"foreignKey:SellerID" json:"products"`
}

type Seller struct {
	Name  string `gorm:"column:name" json:"name"`
	Email string `gorm:"column:email" json:"email"`
}
