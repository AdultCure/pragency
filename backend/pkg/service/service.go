package service

import (
	backend "pragency"
	"pragency/pkg/repository"
)

type Authorization interface {
	CreateUser(user backend.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Order interface {
	Create(UserId int, order backend.Order) (int, error)
	GetAll(UserId int) ([]backend.Order, error)
	GetById(UserId, orderId int) (backend.Order, error)
	Delete(userId, orderId int) error
	Update(userId, orderId int, input backend.UpdateOrderInput) error
}

type Service struct {
	Authorization
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Order:         NewOrderService(repos.Order),
	}
}
