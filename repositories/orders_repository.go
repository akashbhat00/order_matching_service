package repositories

import (
	"order_matching_service/dto"
	"order_matching_service/models"

	"gorm.io/gorm"
)

// OrderRepository defines methods for database operations related to orders.
type OrderRepository interface {
	PlaceOrder(order *dto.CreateOrderRequestDto) (*models.Orders, error)
}

// NewOrderRepository creates a new instance of OrderRepository.
func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepositoryImpl{
		DB: db,
	}
}

type orderRepositoryImpl struct {
	DB *gorm.DB
}

// PlaceOrder places an order by a buyer and returns the matched product.
func (r *orderRepositoryImpl) PlaceOrder(order *dto.CreateOrderRequestDto) (*models.Orders, error) {
	// Find the product with the minimum price, oldest created date, and sufficient quantity
	var matchedProduct models.Products
	if err := r.DB.
		Where("name = ?", order.ProductName).
		Where("quantity >= ?", order.ProductQuantity).
		Order("price ASC, created_at ASC").
		First(&matchedProduct).
		Error; err != nil {
		return nil, err
	}

	// Calculate the updated quantity of the selected product
	updatedQuantity := matchedProduct.Quantity - order.ProductQuantity

	// Create an order with the selected product's ID
	newOrder := &models.Orders{
		BuyerID:    order.BuyerId,
		ProductID:  matchedProduct.ID,
		Quantity:   order.ProductQuantity,
		TotalPrice: float64(order.ProductQuantity) * matchedProduct.Price,
	}

	// Start a database transaction
	tx := r.DB.Begin()

	// Update the quantity of the selected product in the database
	if err := tx.Model(&matchedProduct).Update("quantity", updatedQuantity).Error; err != nil {
		tx.Rollback() // Rollback the transaction if there's an error
		return nil, err
	}

	// Create the order in the database
	if err := tx.Create(newOrder).Error; err != nil {
		tx.Rollback() // Rollback the transaction if there's an error
		return nil, err
	}

	// Commit the transaction
	tx.Commit()

	// Return the selected product
	return newOrder, nil
}
