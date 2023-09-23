package controllers

import (
	"net/http"

	"order_matching_service/dto"
	"order_matching_service/services"

	"github.com/gin-gonic/gin"
)

// OrderController handles HTTP requests related to orders.
type OrderController struct {
	OrderService services.OrderService
}

// NewOrderController creates a new OrderController with the provided OrderService.
func NewOrderController(orderService services.OrderService) *OrderController {
	return &OrderController{
		OrderService: orderService,
	}
}

// PlaceOrder handles the placement of an order by a buyer.
func (ctrl *OrderController) PlaceOrder(c *gin.Context) {
	var order dto.CreateOrderRequestDto
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Call the service to place the order and perform matching
	matchedProduct, err := ctrl.OrderService.PlaceOrder(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to place order"})
		return
	}

	// Return the matched product in the API response
	c.JSON(http.StatusCreated, gin.H{"message": "Order placed", "matched_product": matchedProduct})
}
