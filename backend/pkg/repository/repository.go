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

}

type Repository struct {
	Authorization
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
