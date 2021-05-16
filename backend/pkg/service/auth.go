package service

import (
	"crypto/sha1"
	"fmt"
	backend "pragency"
	"pragency/pkg/repository"
)

const salt = "vlb25hv2lk523j5l23bv5l23"

type AuthServise struct {
	repo repository.Authorization
}

func NewAuthServise(repo repository.Authorization) *AuthServise {
	return &AuthServise{repo: repo}
}

func (s *AuthServise) CreateUser(user backend.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}