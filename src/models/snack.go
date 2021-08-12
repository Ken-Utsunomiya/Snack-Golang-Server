package models

import (
	"time"
)

type Snack struct {
	ID		int		`json:"snack_id"`
	SnackTypeID	int       	`json:"snack_type_id"`
	Name		string 	 	`json:"snack_name"`
	Description	string    	`json:"description"`
	ImageURI	string    	`json:"image_uri"`
	Price		int 		`json:"price"`
	IsActive       	bool 	   	`json:"is_active"`
	OrderThreshold 	int 		`json:"order_threshold"`
	LastUpdatedDTM 	time.Time 	`json:"last_updated_dtm"`
	LastUpdatedBy  	string    	`json:"last_updated_by"`
}

type PopularSnack struct {
	SnackName 	string 	`json:"snack_name"`
	TotalQuantity int    `json:"total_quantity"`
	SnackTypeID   int    `json:"snack_type_id"`
	ImageURI      string `json:"image_uri"`
}

func (Snack) TableName() string {
	return "snacks"
}
