package models

import (
	"time"
)

type Transaction struct {
	ID 						 		uint 			`gorm:"column:transaction_id;primaryKey;not null"`
	SnackName 				string 		`gorm:"column:snack_name;not null;size:128"`
	TransactionAmount uint 			`gorm:"column:transaction_amount;not null"`
	Quantity 					uint 			`gorm:"column:quantity;not null"`
	TransactionDTM 		time.Time `gorm:"column:transaction_dtm;not null"`
}

func (Transaction) TableName() string {
	return "transactions"
}
