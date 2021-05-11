package service

import "pragency/pkg/repository"

type Authorization interface {

}

type Order interface {

}

type Service struct {
	Authorization
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}