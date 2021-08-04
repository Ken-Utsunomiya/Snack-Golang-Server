package models

import (
	"time"
)

type Transaction struct {
	ID			int 		`json:"transaction_id"      gorm:"column:transaction_id;primaryKey;not null"`
	UserID            	int       	`json:"user_id"             gorm:"column:user_id;not null;TYPE:integer REFERENCES users"`
	PaymentID         	int       	`json:"payment_id"          gorm:"column:payment_id;not null;TYPE:integer REFERENCES payments"`
	TransactionTypeID 	int       	`json:"transaction_type_id" gorm:"column:transaction_type_id;not null;TYPE:integer REFERENCES transaction_types"`
	SnackName 		string 		`json:"snack_name"          gorm:"column:snack_name;not null;size:128"`
	TransactionAmount 	int 		`json:"transaction_amount"  gorm:"column:transaction_amount;not null"`
	Quantity 		int 		`json:"quantity"            gorm:"column:quantity;not null"`
	TransactionDTM 		time.Time 	`json:"transaction_dtm"     gorm:"column:transaction_dtm;not null"`
}

func (Transaction) TableName() string {
	return "transactions"
}
