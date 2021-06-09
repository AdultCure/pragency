package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	backend "pragency"
)

// Структура авторизации в бд

type AuthPostgres struct {
	db *sqlx.DB
}

// Функция передачи в дб

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// Функция создания пользователя в бд

func (r *AuthPostgres) CreateUser(user backend.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, email, password_hash, admin) values ($1, $2, $3, $4) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Email, user.Password, user.Admin=="not")
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// Получение пользователя из бд

func (r *AuthPostgres) GetUser(email, password string) (backend.User, error) {
	var user backend.User
	query := fmt.Sprintf("SELECT id, name, admin FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, email, password)

	return user, err
}