package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	backend "pragency"
	"strings"
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
	createOrderQuery := fmt.Sprintf("INSERT INTO %s (category, status, date, comment, user_id, user_name) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", ordersTable)
	row := tx.QueryRow(createOrderQuery, order.Category, order.Status, order.Date, order.Comment, order.UserId, order.UserName)
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

func (r *OrderPostgres) GetAll(userId int) ([]backend.Order, error) {
	var orders []backend.Order

	query := fmt.Sprintf("SELECT tl.id, tl.category, tl. status, tl.date, tl.comment, tl.user_id, tl.user_name FROM %s tl INNER JOIN %s ul on tl.id = ul.order_id WHERE ul.user_id = $1",
		ordersTable, usersListTable)
	err := r.db.Select(&orders, query, userId)

	return orders, err
}

func (r *OrderPostgres) GetAllAdmin() ([]backend.Order, error) {
	var orders []backend.Order

	query := fmt.Sprintf("SELECT tl.id, tl.category, tl. status, tl.date, tl.comment, tl.user_id, tl.user_name FROM %s tl", ordersTable)
	err := r.db.Select(&orders, query)

	return orders, err
}

func (r *OrderPostgres) GetById(userId, orderId int) (backend.Order, error) {
	var order backend.Order

	query := fmt.Sprintf(`SELECT tl.id, tl.category, tl. status, tl.date, tl.comment, tl.user_id, tl.user_name FROM %s tl
								 INNER JOIN %s ul on tl.id = ul.order_id WHERE ul.user_id = $1 AND ul.order_id = $2`,
		ordersTable, usersListTable)
	err := r.db.Get(&order, query, userId, orderId)

	return order, err
}

func (r *OrderPostgres) Delete(userId, orderId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.order_id AND ul.user_id=$1 AND ul.order_id=$2",
		ordersTable, usersListTable)
	_, err := r.db.Exec(query, userId, orderId)

	return err
}

func (r *OrderPostgres) Update(userId, orderId int, input backend.UpdateOrderInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Category != nil {
		setValues = append(setValues, fmt.Sprintf("category=$%d", argId))
		args  = append(args, *input.Category)
		argId++
	}

	if input.Status != nil {
		setValues = append(setValues, fmt.Sprintf("status=$%d", argId))
		args  = append(args, *input.Status)
		argId++
	}

	if input.Date != nil {
		setValues = append(setValues, fmt.Sprintf("date=$%d", argId))
		args  = append(args, *input.Date)
		argId++
	}

	if input.Comment != nil {
		setValues = append(setValues, fmt.Sprintf("comment=$%d", argId))
		args  = append(args, *input.Comment)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.order_id AND ul.order_id=$%d AND ul.user_id=$%d",
		ordersTable, setQuery, usersListTable, argId, argId+1)
	args = append(args, orderId, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}