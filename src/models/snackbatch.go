package models

import (
	"database/sql"
)

type SnackBatch struct {
	ID	int	`json:"snack_batch_id"`
	SnackID	int	`json:"snack_id"`
	Quantity	int	`json:"quantity"`
	ExpirationDTM	sql.NullTime	`json:"expiration_dtm"`
}

func (SnackBatch) TableName() string {
	return "snack_batches"
}
