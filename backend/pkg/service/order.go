package service

import (
	backend "pragency"
	"pragency/pkg/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(userId int, order backend.Order) (int, error) {
	return s.repo.Create(userId, order)
}

func (s  *OrderService) GetAll(userId  int) ([]backend.Order, error) {
	return s.repo.GetAll(userId)
}

func (s *OrderService) GetAllAdmin() ([]backend.Order, error) {
	return s.repo.GetAllAdmin()
}

func (s *OrderService) GetById(userId, orderId int) (backend.Order, error) {
	return s.repo.GetById(userId, orderId)
}

func (s *OrderService) Delete(userId, orderId int) error {
	return s.repo.Delete(userId, orderId)
}

func (s *OrderService) Update(userId, orderId int, input backend.UpdateOrderInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, orderId, input)
}