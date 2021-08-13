package models

import (
	"time"
)

type User struct {
	ID		int 	 	`json:"user_id"       gorm:"column:user_id;primaryKey;not null"`
	Email		string 		`json:"email_address" gorm:"column:email_address;unique;not null;size:128"`
	FirstName 	string 		`json:"first_name"    gorm:"column:first_name;not null;size:128"`
	LastName 	string 		`json:"last_name"     gorm:"column:last_name;not null;size:128"`
	ImageURI 	string		`json:"image_uri"     gorm:"column:image_uri;not null"`
	Balance 	int 	 	`json:"balance"       gorm:"column:balance;not null;default:0"`
	IsAdmin 	bool 	 	`json:"is_admin"      gorm:"column:is_admin;not null;default:false"`
	IsActive 	bool		`json:"is_active"     gorm:"column:is_active;not null;default:true"`
	DeletedAt time.Time 	`json:"deleted_at"    gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "users"
}
