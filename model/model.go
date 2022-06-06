package model

import "time"

type Transaction struct {
	ID            uint   `json:"-" gorm:"primaryKey"`
	IdTransaction uint   `json:"id_transaction" gorm:"primaryKey;autoIncrement"`
	IdUser        int    `json:"id_user" gorm:"not null"`
	Email         string `json:"email" `
	Price         int    `json:"price" `
	PriceType     string `json:"price_type" `
	Status        string `json:"status" `
	CreateAt      time.Time
	UpdatedAt     time.Time
}

type PaySystem struct {
	PaySystemId int `json:"pay_system_id"`
}

type DataTransaction struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}
