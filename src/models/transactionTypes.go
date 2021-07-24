package models

type TransactionType struct {
	ID 	 uint 	`gorm:"column:transaction_type_id;primaryKey;not null"`
	Name string `json:"column:transaction_type_name;unique;not null;size:128"`
}

func (TransactionType) TableName() string {
	return "transaction_types"
}
