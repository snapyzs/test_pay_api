package service

import (
	"test_project_sell/internal/repository"
	"test_project_sell/model"
)

type Auth interface {
	GenerateToken(psi int) (string, error)
	ParseToken(tokenAccess string) (int, error)
}

type Transaction interface {
	CreatePay(trans model.Transaction) error
	EditStatusPay(trans model.Transaction) error
	GetAllPayUserById(trans model.Transaction) ([]model.Transaction, error)
	GetAllPayUserByEmail(trans model.Transaction) ([]model.Transaction, error)
	CheckStatus(trans model.Transaction) (string, error)
}

type Service struct {
	Auth
	Transaction
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Auth:        NewAuthService(r.Auth),
		Transaction: NewTransactionService(r.Transaction),
	}
}
