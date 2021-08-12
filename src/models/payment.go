package models

import (
	"time"
)

type Payment struct {
	ID		int		`json:"payment_id"`
	UserID        	int       	`json:"user_id"`
	PaymentAmount 	int 		`json:"payment_amount"`
	PaymentDTM 	time.Time 	`json:"payment_dtm"`
	CreatedBy 	string 		`json:"created_by"`
}

func (Payment) TableName() string {
	return "payments"
}
