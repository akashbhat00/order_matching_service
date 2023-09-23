package services

import (
	"order_matching_service/dto"
	"order_matching_service/models"
	"order_matching_service/repositories"
)

// OrderService defines methods related to order operations.
type OrderService interface {
	PlaceOrder(order *dto.CreateOrderRequestDto) (*models.Orders, error)
}

// NewOrderService creates a new instance of OrderService.
func NewOrderService(orderRepo repositories.OrderRepository) OrderService {
	return &orderServiceImpl{
		orderRepo: orderRepo,
	}
}

type orderServiceImpl struct {
	orderRepo repositories.OrderRepository
}

// PlaceOrder places an order by a buyer and returns the matched product.
func (s *orderServiceImpl) PlaceOrder(order *dto.CreateOrderRequestDto) (*models.Orders, error) {
	// Call the repository to place the order and perform matching
	orderResponse, err := s.orderRepo.PlaceOrder(order)
	if err != nil {
		return nil, err
	}

	return orderResponse, nil
}
