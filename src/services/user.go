package services

import (
	"Snack-Golang-Server/src/database"
	"Snack-Golang-Server/src/models"
)

type UserService struct {}

func (UserService) GetUserList() []models.User {
	return nil
}

func (UserService) GetUser(Id int) (models.User, error) {
	db := database.GetDB()
	user := models.User{}
	err := db.First(&user, Id).Error
	return user, err
}

func (UserService) GetUserCommonList() ([]models.User, error) {
	db := database.GetDB()
	users := make([]models.User, 0)
	err := db.Find(&users).Error
	return users, err
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
