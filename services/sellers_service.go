// services/seller_service.go
package services

import (
	"order_matching_service/dto"
	"order_matching_service/models"
	"order_matching_service/repositories"
)

type SellerService struct {
	SellerRepo *repositories.SellerRepository
}

func NewSellerService(sellerRepo *repositories.SellerRepository) *SellerService {
	return &SellerService{SellerRepo: sellerRepo}
}

// CreateSeller creates a new seller using the repository.
func (s *SellerService) CreateSeller(seller *models.Sellers) error {
	return s.SellerRepo.CreateSeller(seller)
}

func (s *SellerService) GetSellerByID(sellerID uint) (*dto.SellerWithProducts, error) {
	return s.SellerRepo.GetSellerByID(sellerID)
}

func (s *SellerService) AddProductToCatalog(sellerID uint, product *dto.ProductRequestDto) (*models.Products, error) {
	// Call the repository to perform the database operation
	return s.SellerRepo.AddProductToCatalog(sellerID, product)
}
