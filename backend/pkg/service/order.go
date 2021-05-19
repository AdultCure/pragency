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