package repository

import (
	"gorm.io/gorm"
	"test_project_sell/model"
)

type Auth interface {
}

type Transaction interface {
	CreatePay(trans model.Transaction) error
	EditStatusPay(trans model.Transaction) error
	GetAllPayUserById(trans model.Transaction) ([]model.Transaction, error)
	GetAllPayUserByEmail(trans model.Transaction) ([]model.Transaction, error)
	CheckStatus(trans model.Transaction) (string, error)
}

type Repository struct {
	Auth
	Transaction
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Auth:        NewAuthPostgress(db),
		Transaction: NewTransactionPostgres(db),
	}
}
