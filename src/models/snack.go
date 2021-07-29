package models

import (
	"time"
)

type Snack struct {
	ID 			       uint 		 `gorm:"column:snack_id;primaryKey;autoIncrement"`
	SnackTypeID    uint      `gorm:"column:snack_type_id;not null"`
	Name           string 	 `gorm:"column:snack_name;unique;not null;size:128"`
	Description    string    `gorm:"column:description;not null;size:128"`
	ImageURI       string    `gorm:"column:image_uri;not null"`
	Price          uint 		 `gorm:"column:price;not null"`
	IsActive       bool 	   `gorm:"column:is_active;not null"`
	OrderThreshold uint 		 `gorm:"column:order_threshold"`
	LastUpdatedDTM time.Time `gorm:"column:last_updated_dtm;not null"`
	LastUpdatedBy  string    `gorm:"column:last_updated_by;not null;size:128"`
	SnackType      SnackType
}

func (Snack) TableName() string {
	return "snacks"
}
