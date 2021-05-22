package repository

import (
	"github.com/jmoiron/sqlx"
	backend "pragency"
)

type Authorization interface {
	CreateUser(user backend.User) (int, error)
	GetUser(email, password string) (backend.User, error)
}

type Order interface {
	Create(userId int, order backend.Order) (int, error)
	GetAll(userId int ) ([]backend.Order, error)
	GetById(userId, orderId int) (backend.Order, error)
	Delete(userId, orderId int) error
	Update(userId, orderId int, input backend.UpdateOrderInput) error
}

type Repository struct {
	Authorization
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Order: NewOrderPostgres(db),
	}
}
