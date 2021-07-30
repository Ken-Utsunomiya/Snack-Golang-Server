package services

import "Snack-Golang-Server/src/models"

type UserService struct {}

func (UserService) GetUserList() []models.User {
	return nil
}

func (UserService) GetUser() *models.User {
	return nil
}

func (UserService) GetUserCommonList() []models.User {
	return nil
}

func (UserService) AddUser(user *models.User) error {
	return nil
}

func (UserService) UpdateUser(user *models.User) error {
	return nil
}

func (UserService) DeleteUser(id uint) error {
	return nil
}
