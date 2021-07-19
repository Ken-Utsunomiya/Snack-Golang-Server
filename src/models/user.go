package models

type User struct {
	Id int `json:"userId"`
	Email string `json:"email"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	ImageURI string `json:"imageURI"`
	Balance int `json:"balance"`
	IsAdmin bool `json:"isAdmin"`
	IsActive bool `json:"isActive"`
}
