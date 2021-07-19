package models

import "github.com/golang/protobuf/ptypes/timestamp"

type Transaction struct {
	Id int `json:"transactionId"`
	SnackName string `json:"snackName"`
	TransactionAmount int `json:"transactionAmount"`
	Quantity int `json:"quantity"`
	TransactionDTM timestamp.Timestamp `json:"transactionDTM"`
}
