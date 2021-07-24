package models

import (
	"time"
)

type SnackBatch struct {
	ID 						uint 			`gorm:"column:snack_batch_id;primaryKey;not null"`
	Quantity 			uint 			`gorm:"column:quantity;not null"`
	ExpirationDTM time.Time `gorm:"column:expiration_dtm"`
}

func (SnackBatch) TableName() string {
	return "snack_batches"
}
