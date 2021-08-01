package models

type SnackType struct {
	ID 	 int 	  `json:"snack_type_id"   gorm:"column:snack_type_id;primaryKey;not null"`
	Name string `json:"snack_type_name" gorm:"column:snack_type_name;unique;not null;size:128"`
}

func (SnackType) TableName() string {
	return "snack_types"
}
