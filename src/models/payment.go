package models

import (
	"time"
)

type Payment struct {
	ID 						uint 			`json:"payment_id" gorm:"column:payment_id;primaryKey;autoIncrement;not null"`
	UserID        uint      `json:"user_id" gorm:"column:user_id;not null"`
	PaymentAmount uint 			`json:"payment_amount" gorm:"column:payment_amount;not null"`
	PaymentDTM 		time.Time `json:"payment_dtm" gorm:"column:payment_dtm;not null"`
	CreatedBy 		string 		`json:"created_by" gorm:"column:created_by;not null;size:128"`
	User          User
}

func (Payment) TableName() string {
	return "payments"
}
