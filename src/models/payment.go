package models

import "github.com/golang/protobuf/ptypes/timestamp"

type Payment struct {
	PaymentId int `json:"paymentId"`
	PaymentAmount int `json:"paymentAmount"`
	PaymentDTM timestamp.Timestamp `json:"paymentDTM"`
	CreatedBy string `json:"CreatedBy"`
}
