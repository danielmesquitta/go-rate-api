package model

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod string

const (
	CreditCard PaymentMethod = "DEBIT_CARD"
	DebitCard  PaymentMethod = "CREDIT_CARD"
	Cash       PaymentMethod = "CASH"
)

type User struct {
	ID        uint           `json:"id,omitempty" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty" gorm:"index"`
	Name      string         `json:"name,omitempty"`
	Email     string         `json:"email,omitempty" gorm:"unique"`
	Password  string         `json:"password,omitempty"`
}
