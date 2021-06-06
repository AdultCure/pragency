package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	backend "pragency"
	"pragency/pkg/repository"
	"time"
)

// соль для хэширования пароля
const (
	// соль
	salt = "vlb25hv2lk523j5l23bv5l23"
	// ключ подписи
	signingKey = "3k4h5lk34h5l3k4&#@"
	// время валидности токина
	tokenTTL = 128 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user backend.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	// получение пользователя из бд
	user, err := s.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			//валидность токена ограничена 12 часами по условию tokenTTL
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			//время когда токен был сгенерирован
			IssuedAt: time.Now().Unix(),
		}, user.Id,
	})

	//получает набор случайных пайтов для подписи и расшифровки
	return token.SignedString([]byte(signingKey))
}

func  (s *AuthService) GetUserMainParams(email, password string) (string, int, string, error) {
	user, err := s.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", 0, "", nil
	}

	name := user.Name
	id := user.Id
	admin := user.Admin

	return name, id, admin, nil
}

func (s *AuthService) GetAdminRole(email, password string) (string, error) {
	user, err := s.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	admin := user.Admin

	return admin, nil
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

// хэширование пароля

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
