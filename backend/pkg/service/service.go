package service

import (
	backend "pragency"
	"pragency/pkg/repository"
)

// Интерфейс с функция авторизации

type Authorization interface {
	CreateUser(user backend.User) (int, error)
	GenerateToken(email, password string) (string, error)
	GetUserMainParams(email, password string) (string, int, string, error)
	ParseToken(token string) (int, error)
}

// Интерфейс с функция заказов

type Order interface {
	Create(UserId int, order backend.Order) (int, error)
	GetAll(UserId int) ([]backend.Order, error)
	GetAllAdmin() ([]backend.Order, error)
	GetOneAdmin(orderId int) (backend.Order, error)
	UpdateAdmin(orderId int, input backend.UpdateOrderInput) error
	DeleteAdmin(orderId int) error
	GetById(UserId, orderId int) (backend.Order, error)
}

// Структура сервисов

type Service struct {
	Authorization
	Order
}

//

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Order:         NewOrderService(repos.Order),
	}
}
