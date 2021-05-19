package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	backend "pragency"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) Create(userId int, order backend.Order) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createOrderQuery := fmt.Sprintf("INSERT INTO %s (category, status, date, comment) VALUES ($1, $2, $3, $4) RETURNING id", ordersTable)
	row := tx.QueryRow(createOrderQuery, order.Category, order.Status, order.Date, order.Comment)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, order_id) VALUES ($1, $2)", usersListTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}