package models

import (
	"time"
)

type User struct {
	ID		int 	 	`json:"user_id"`
	Email		string 		`json:"email_address"`
	FirstName 	string 		`json:"first_name"`
	LastName 	string 		`json:"last_name"`
	ImageURI 	string		`json:"image_uri"`
	Balance 	int 	 	`json:"balance"`
	IsAdmin 	bool 	 	`json:"is_admin"`
	IsActive 	bool		`json:"is_active"`
	DeletedAt 	time.Time 	`json:"deleted_at"`
}

func (User) TableName() string {
	return "users"
}
