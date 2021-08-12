package models

type SnackType struct {
	ID	int	`json:"snack_type_id"`
	Name 	string	`json:"snack_type_name"`
}

func (SnackType) TableName() string {
	return "snack_types"
}
