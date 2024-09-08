package usecases

import "github.com/thanchayawikgithub/learn-go-clean/entities"

type OrderUseCase interface {
	CreateOrder(order entities.Order) error
}

type OrderService struct {
	orderRepo OrderRepository
}

func NewOrderService(orderRepo OrderRepository) OrderUseCase {
	return &OrderService{orderRepo}
}

func (s *OrderService) CreateOrder(order entities.Order) error {
	return s.orderRepo.Save(order)
}
