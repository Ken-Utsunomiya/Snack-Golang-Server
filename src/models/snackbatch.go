package models

import (
	"time"
)

type SnackBatch struct {
	ID 						int 			`json:"snack_batch_id" gorm:"column:snack_batch_id;primaryKey;not null"`
	SnackID       int       `json:"snack_id"       gorm:"column:snack_id;not null"`
	Quantity 			int 			`json:"quantity"       gorm:"column:quantity;not null"`
	ExpirationDTM time.Time `json:"expiration_dtm" gorm:"column:expiration_dtm"`
	Snack         Snack
}

func (SnackBatch) TableName() string {
	return "snack_batches"
}
