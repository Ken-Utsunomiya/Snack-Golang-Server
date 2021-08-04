package models

import (
	"database/sql"
)

type SnackBatch struct {
	ID 		int		`json:"snack_batch_id" gorm:"column:snack_batch_id;primaryKey;not null"`
	SnackID       	int          	`json:"snack_id"       gorm:"column:snack_id;not null;TYPE:integer REFERENCES snacks"`
	Quantity 	int 		`json:"quantity"       gorm:"column:quantity;not null"`
	ExpirationDTM 	sql.NullTime 	`json:"expiration_dtm" gorm:"column:expiration_dtm"`
}

func (SnackBatch) TableName() string {
	return "snack_batches"
}
