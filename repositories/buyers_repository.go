// repositories/seller_repository.go
package repositories

import (
	"order_matching_service/models"

	"gorm.io/gorm"
)

type BuyerRepository struct {
	DB *gorm.DB
}

func NewBuyerRepository(db *gorm.DB) *BuyerRepository {
	return &BuyerRepository{DB: db}
}

// CreateSeller creates a new seller in the database.
func (repo *BuyerRepository) CreateBuyer(buyer *models.Buyers) error {
	return repo.DB.Create(buyer).Error
}

func (repo *BuyerRepository) GetBuyerByID(buyerID uint) (*models.Buyers, error) {
	var buyer models.Buyers

	// Find the seller by ID
	if err := repo.DB.First(&buyer, buyerID).Error; err != nil {
        return nil, err
    }

	return &buyer, nil
}
