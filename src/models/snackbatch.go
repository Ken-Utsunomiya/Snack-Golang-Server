package models

import "github.com/golang/protobuf/ptypes/timestamp"

type SnackBatch struct {
	SnackBatchId int `json:"snackBatchId"`
	Quantity int `json:"quantity"`
	ExpirationDTM timestamp.Timestamp `json:"expirationDTM"`
}
