package models

import "time"

type User struct {
	ID 				uint 	 		`gorm:"column:user_id;primaryKey;not null"`
	Email 		string 		`gorm:"column:email_address;unique;not null;size:128"`
	FirstName string 		`gorm:"column:first_name;not null;size:128"`
	LastName 	string 		`gorm:"column:last_name;not null;size:128"`
	ImageURI 	string 		`gorm:"column:image_uri;not null"`
	Balance 	int 	 		`gorm:"column:balance;not null"`
	IsAdmin 	bool 	 		`gorm:"column:is_admin;not null"`
	IsActive 	bool 	 		`gorm:"column:is_active;not null"`
	DeletedAt time.Time `gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "users"
}
