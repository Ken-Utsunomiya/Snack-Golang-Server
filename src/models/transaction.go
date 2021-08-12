package models

import (
	"time"
)

type Transaction struct {
	ID			int 		`json:"transaction_id"`
	UserID            	int     	`json:"user_id"`
	PaymentID         	int     	`json:"payment_id"`
	TransactionTypeID 	int     	`json:"transaction_type_id"`
	SnackName 		string 		`json:"snack_name"`
	TransactionAmount 	int 		`json:"transaction_amount"`
	Quantity 		int 		`json:"quantity"`
	TransactionDTM 		time.Time 	`json:"transaction_dtm"`
}

func (Transaction) TableName() string {
	return "transactions"
}
