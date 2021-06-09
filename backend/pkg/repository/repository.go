package repository

import (
	"github.com/jmoiron/sqlx"
	backend "pragency"
)

// Интерфейс авторизации

type Authorization interface {
	CreateUser(user backend.User) (int, error)
	GetUser(email, password string) (backend.User, error)
}

// Интерфейс заказов

type Order interface {
	Create(userId int, order backend.Order) (int, error)
	GetAll(userId int ) ([]backend.Order, error)
	GetAllAdmin() ([]backend.Order, error)
	GetOneAdmin(orderId int) (backend.Order, error)
	UpdateAdmin(orderId int, input backend.UpdateOrderInput) error
	DeleteAdmin(orderId int) error
	GetById(userId, orderId int) (backend.Order, error)
}

// Структура репозитория

type Repository struct {
	Authorization
	Order
}

// Функция создания нового репозитория

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Order: NewOrderPostgres(db),
	}
}
