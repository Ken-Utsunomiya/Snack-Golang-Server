package models

import (
	"time"
)

type Snack struct {
	ID		int		`json:"snack_id" gorm:"column:snack_id;primaryKey;autoIncrement"`
	SnackTypeID	int       	`json:"snack_type_id" gorm:"column:snack_type_id;not null;"`
	Name		string 	 	`json:"snack_name" gorm:"column:snack_name;unique;not null;size:128"`
	Description	string    	`json:"description" gorm:"column:description;not null;size:128"`
	ImageURI	string    	`json:"image_uri" gorm:"column:image_uri;not null"`
	Price		int 		`json:"price" gorm:"column:price;not null"`
	IsActive       	bool 	   	`json:"is_active" gorm:"column:is_active;not null"`
	OrderThreshold 	*int 		`json:"order_threshold" gorm:"column:order_threshold"`
	LastUpdatedDTM 	time.Time 	`json:"last_updated_dtm" gorm:"column:last_updated_dtm;not null"`
	LastUpdatedBy  	string    	`json:"last_updated_by" gorm:"column:last_updated_by;not null;size:128"`
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
