package repository

import "gorm.io/gorm"

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgress(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
