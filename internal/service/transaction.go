package service

import (
	"test_project_sell/internal/repository"
	"test_project_sell/model"
)

type TransactionService struct {
	repository repository.Transaction
}

func NewTransactionService(r repository.Transaction) *TransactionService {
	return &TransactionService{
		repository: r,
	}
}

func (t *TransactionService) CreatePay(trans model.Transaction) error {
	return t.repository.CreatePay(trans)
}

func (t *TransactionService) EditStatusPay(trans model.Transaction) error {
	return t.repository.EditStatusPay(trans)
}

func (t *TransactionService) GetAllPayUserById(trans model.Transaction) ([]model.Transaction, error) {
	return t.repository.GetAllPayUserById(trans)
}

func (t *TransactionService) GetAllPayUserByEmail(trans model.Transaction) ([]model.Transaction, error) {
	return t.repository.GetAllPayUserByEmail(trans)
}

func (t *TransactionService) CheckStatus(trans model.Transaction) (string, error) {
	return t.repository.CheckStatus(trans)
}
