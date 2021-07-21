package models

import (
	"time"
)

type Payment struct {
	PaymentId int `json:"paymentId" gorm:"primaryKey;autoIncrement;not null"`
	PaymentAmount int `json:"paymentAmount" gorm:"not null"`
	PaymentDTM time.Time `json:"paymentDTM" gorm:"not null"`
	CreatedBy string `json:"CreatedBy" gorm:"not null"`
}
