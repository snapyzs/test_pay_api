package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"test_project_sell/internal/repository"
	"time"
)

const (
	signingKey = "qwpoxcbzxui"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	PaySystemId int `json:"pay_system_id"`
}

type AuthService struct {
	repository repository.Auth
}

func NewAuthService(r repository.Auth) *AuthService {
	return &AuthService{repository: r}
}

func (a *AuthService) GenerateToken(paySystemId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		}, paySystemId})
	return token.SignedString([]byte(signingKey))
}

func (a *AuthService) ParseToken(tokenAccess string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenAccess, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
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
	return claims.PaySystemId, nil
}
