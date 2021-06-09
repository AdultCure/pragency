package backend

import "errors"

// Структура заказа

type Order struct {
	Id       int    `json:"id" db:"id"`
	Category string `json:"category" db:"category" binding:"required"`
	Status   string `json:"status" db:"status" binding:"required"`
	Date     string `json:"date" db:"date" binding:"required"`
	Comment  string `json:"comment" db:"comment" binding:"required"`
	UserId   string `json:"user_id" db:"user_id" binding:"required"`
	UserName string `json:"user_name" db:"user_name" binding:"required"`
}

// Структура для обновления заказа

type UpdateOrderInput struct {
	Category *string `json:"category"`
	Status   *string `json:"status"`
	Date     *string `json:"date"`
	Comment  *string `json:"comment"`
}

// Функция валидации обновления заказа

func (i UpdateOrderInput) Validate() error {
	if i.Category == nil && i.Status == nil && i.Date == nil && i.Comment == nil {
		return errors.New("update structure has no values")
	}

	return nil
}