package models

type TransactionType struct {
	ID 	 uint 	`json:"transaction_type_id"   gorm:"column:transaction_type_id;primaryKey;not null"`
	Name string `json:"transaction_type_name" gorm:"column:transaction_type_name;unique;not null;size:128"`
}

func (TransactionType) TableName() string {
	return "transaction_types"
}
