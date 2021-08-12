package models

type TransactionType struct {
	ID	int	`json:"transaction_type_id"`
	Name 	string 	`json:"transaction_type_name"`
}

func (TransactionType) TableName() string {
	return "transaction_types"
}
