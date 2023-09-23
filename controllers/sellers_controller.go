// controllers/seller_controller.go
package controllers

import (
	"net/http"
	"strconv"

	"order_matching_service/dto"
	"order_matching_service/models"
	"order_matching_service/services"

	"github.com/gin-gonic/gin"
)

type SellerController struct {
	SellerService *services.SellerService
}

func NewSellerController(sellerService *services.SellerService) *SellerController {
	return &SellerController{SellerService: sellerService}
}

// CreateSeller creates a new seller profile.
func (ctrl *SellerController) CreateSeller(c *gin.Context) {
	var requestDTO dto.CreateSellerRequest

	// Bind the request body to the request DTO
	if err := c.ShouldBindJSON(&requestDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create a new seller model using the request data
	seller := models.Sellers{
		Name:  requestDTO.Name,
		Email: requestDTO.Email,
		// Set other fields as needed
	}

	// Call the seller service to create the seller
	if err := ctrl.SellerService.CreateSeller(&seller); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create seller"})
		return
	}

	// Return a success response
	c.JSON(http.StatusCreated, gin.H{"data": seller})
}

func (ctrl *SellerController) GetSellerByID(c *gin.Context) {
	// Get the seller ID from the URL parameter
	sellerIDParam := c.Param("seller_id")
	sellerID, err := strconv.ParseUint(sellerIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seller ID"})
		return
	}

	// Call the service to retrieve the seller
	seller, err := ctrl.SellerService.GetSellerByID(uint(sellerID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Seller not found"})
		return
	}

	// Return the seller as JSON response
	c.JSON(http.StatusOK, gin.H{"data": seller})
}

func (ctrl *SellerController) AddProductToSellerCatalog(c *gin.Context) {
	sellerIDParam := c.Param("seller_id")
	sellerID, err := strconv.ParseUint(sellerIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seller ID"})
		return
	}

	var product dto.ProductRequestDto
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Call the service to add the product to the seller's catalog
	response, err := ctrl.SellerService.AddProductToCatalog(uint(sellerID), &product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product added to seller's catalog", "data": response})
}
