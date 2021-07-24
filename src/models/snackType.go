package models

type SnackType struct {
	ID 	 uint 	`gorm:"column:snack_type_id;not null"`
	Name string `gorm:"column:snack_type_name;not null;size:128"`
}

func (SnackType) TableName() string {
	return "snack_types"
}
