package repository

import (
	"errors"
	"gorm.io/gorm"
	"test_project_sell/model"
)

type TransactionPostgres struct {
	db *gorm.DB
}

func NewTransactionPostgres(db *gorm.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (t *TransactionPostgres) CreatePay(trans model.Transaction) error {
	tx := t.db.Begin()
	if err := tx.Create(&trans).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}
func (t *TransactionPostgres) EditStatusPay(trans model.Transaction) error {
	type Transactions struct {
		IdTransaction int
	}
	tx := t.db.Begin()
	err := tx.First(&Transactions{}, trans.IdTransaction).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if err := tx.Model(&model.Transaction{}).Where("id_transaction = ?", trans.IdTransaction).Update("status", trans.Status).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (t *TransactionPostgres) GetAllPayUserById(trans model.Transaction) ([]model.Transaction, error) {
	var tran []model.Transaction
	tx := t.db.Begin()
	if err := tx.Model(&model.Transaction{}).Where("id_user = ?", trans.IdUser).Find(&tran).Error; err != nil {
		return nil, err
	}
	return tran, nil
}
func (t *TransactionPostgres) GetAllPayUserByEmail(trans model.Transaction) ([]model.Transaction, error) {
	var tran []model.Transaction
	tx := t.db.Begin()
	if err := tx.Model(&model.Transaction{}).Where("email = ?", trans.Email).Find(&tran).Error; err != nil {
		return nil, err
	}
	return tran, nil
}

func (t *TransactionPostgres) CheckStatus(trans model.Transaction) (string, error) {
	tx := t.db.Begin()
	if err := tx.Model(&model.Transaction{}).Where("id_transaction = ?", trans.IdTransaction).First(&trans).Error; err != nil {
		return "", err
	}
	return trans.Status, nil
}
