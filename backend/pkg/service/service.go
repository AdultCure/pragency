package service

import (
	backend "pragency"
	"pragency/pkg/repository"
)

type Authorization interface {
	CreateUser(user backend.User) (int, error)
	GenerateToken(email, password string) (string, error)
}

type Order interface {

}

type Service struct {
	Authorization
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}