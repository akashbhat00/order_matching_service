package dto

type CreateOrderRequestDto struct {
	BuyerId         uint   `json:"buyer_id"`
	ProductName     string `json:"product_name"`
	ProductQuantity uint   `json:"product_quantity"`
}
