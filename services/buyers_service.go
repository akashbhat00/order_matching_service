package services

import (
	"order_matching_service/models"
	"order_matching_service/repositories"
)

type BuyerService struct {
	BuyerRepo *repositories.BuyerRepository
}

func NewBuyerService(buyerRepo *repositories.BuyerRepository) *BuyerService {
	return &BuyerService{BuyerRepo: buyerRepo}
}

// CreateBuyer creates a new buyer using the repository.
func (s *BuyerService) CreateBuyer(buyer *models.Buyers) error {
	return s.BuyerRepo.CreateBuyer(buyer)
}

func (s *BuyerService) GetBuyerByID(buyerID uint) (*models.Buyers, error) {
	return s.BuyerRepo.GetBuyerByID(buyerID)
}
