package models

import "github.com/golang/protobuf/ptypes/timestamp"

type Snack struct {
	SnackId int `json:"snackId"`
	SnackName string `json:"snackName"`
	Description string `json:"description"`
	ImageURI string `json:"imageURI"`
	Price int `json:"price"`
	IsActive bool `json:"isActive"`
	OrderThreshold int `json:"orderThreshold"`
	LastUpdatedDTM timestamp.Timestamp `json:"LastUpdatedDTM"`
	LastUpdatedBy string `json:"LastUpdatedBy"`
}
