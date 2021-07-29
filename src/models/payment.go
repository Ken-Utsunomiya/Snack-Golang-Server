package models

import (
	"time"
)

type Payment struct {
	ID 						uint 			`gorm:"column:payment_id;primaryKey;autoIncrement;not null"`
	UserID        uint      `gorm:"column:user_id;not null"`
	PaymentAmount uint 			`gorm:"column:payment_amount;not null"`
	PaymentDTM 		time.Time `gorm:"column:payment_dtm;not null"`
	CreatedBy 		string 		`gorm:"column:created_by;not null;size:128"`
	User          User
}

func (Payment) TableName() string {
	return "payments"
}
