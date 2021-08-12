package models

import (
	"time"
)

type SnackBatch struct {
	ID	int	`json:"snack_batch_id"`
	SnackID	int	`json:"snack_id"`
	Quantity	int	`json:"quantity"`
	ExpirationDTM	time.Time	`json:"expiration_dtm"`
}

func (SnackBatch) TableName() string {
	return "snack_batches"
}
