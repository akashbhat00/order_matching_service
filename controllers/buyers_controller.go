package controllers

import (
	"net/http"
	"strconv"

	"order_matching_service/dto"
	"order_matching_service/models"
	"order_matching_service/services"

	"github.com/gin-gonic/gin"
)

type BuyerController struct {
	BuyerService *services.BuyerService
}

func NewBuyerController(buyerService *services.BuyerService) *BuyerController {
	return &BuyerController{BuyerService: buyerService}
}

// CreateBuyer creates a new buyer profile.
func (ctrl *BuyerController) CreateBuyer(c *gin.Context) {
	var requestDTO dto.CreateBuyerRequest

	// Bind the request body to the request DTO
	if err := c.ShouldBindJSON(&requestDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new seller model using the request data
	buyer := models.Buyers{
		Name:  requestDTO.Name,
		Email: requestDTO.Email,
		// Set other fields as needed
	}

	// Call the seller service to create the seller
	if err := ctrl.BuyerService.CreateBuyer(&buyer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create buyer"})
		return
	}

	// Return a success response
	c.JSON(http.StatusCreated, gin.H{"data": buyer})
}

func (ctrl *BuyerController) GetBuyerByID(c *gin.Context) {
	// Get the seller ID from the URL parameter
	buyerIDParam := c.Param("buyer_id")
	buyerID, err := strconv.ParseUint(buyerIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seller ID"})
		return
	}

	// Call the service to retrieve the seller
	seller, err := ctrl.BuyerService.GetBuyerByID(uint(buyerID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Seller not found"})
		return
	}

	// Return the seller as JSON response
	c.JSON(http.StatusOK, gin.H{"data": seller})
}
