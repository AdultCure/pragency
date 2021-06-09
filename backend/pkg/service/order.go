package service

import (
	backend "pragency"
	"pragency/pkg/repository"
)

// Структура сервиса заказов

type OrderService struct {
	repo repository.Order
}

// Функция создания нового сервича заказа

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

// Функция создания нового заказа

func (s *OrderService) Create(userId int, order backend.Order) (int, error) {
	return s.repo.Create(userId, order)
}

// Функция получение всех заказов у пользователя

func (s  *OrderService) GetAll(userId  int) ([]backend.Order, error) {
	return s.repo.GetAll(userId)
}

// Функция получения всех заказов для админа

func (s *OrderService) GetAllAdmin() ([]backend.Order, error) {
	return s.repo.GetAllAdmin()
}

// Функция получение одного заказа для пользователя

func (s *OrderService) GetById(userId, orderId int) (backend.Order, error) {
	return s.repo.GetById(userId, orderId)
}

// Функция получение одного заказа для админа

func (s *OrderService) GetOneAdmin(orderId int) (backend.Order, error) {
	return s.repo.GetOneAdmin(orderId)
}

// Функция обновления заказа для админа

func (s *OrderService) UpdateAdmin(orderId int, input backend.UpdateOrderInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateAdmin(orderId, input)
}

// Функция удаления заказа для админа

func (s *OrderService) DeleteAdmin(orderId int) error {
	return s.repo.DeleteAdmin(orderId)
}