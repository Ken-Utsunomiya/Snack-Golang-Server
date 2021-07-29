package models

import (
	"time"
)

type Transaction struct {
	ID 						 		uint 			`gorm:"column:transaction_id;primaryKey;not null"`
	UserID            uint      `gorm:"column:user_id;not null"`
	PaymentID         uint      `gorm:"column:payment_id;not null"`
	TransactionTypeID uint      `gorm:"column:transaction_type_id;not null"`
	SnackName 				string 		`gorm:"column:snack_name;not null;size:128"`
	TransactionAmount uint 			`gorm:"column:transaction_amount;not null"`
	Quantity 					uint 			`gorm:"column:quantity;not null"`
	TransactionDTM 		time.Time `gorm:"column:transaction_dtm;not null"`
	User              User
	Payment           Payment
	TransactionType   TransactionType
}

func (Transaction) TableName() string {
	return "transactions"
}
