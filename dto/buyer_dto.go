package dto

// CreateSellerRequest represents the request data for creating a buyer.
type CreateBuyerRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}
