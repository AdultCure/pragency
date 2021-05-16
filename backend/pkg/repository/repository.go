package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {

}

type Order interface {

}

type Repository struct {
	Authorization
	Order
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
