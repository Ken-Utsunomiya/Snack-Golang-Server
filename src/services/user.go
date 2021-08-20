package services

import (
	"Snack-Golang-Server/src/database"
	"Snack-Golang-Server/src/models"
	"Snack-Golang-Server/src/validators"
)

type UserService struct {}

func (UserService) GetUserList(email string) ([]models.User, error) {
	db := database.GetDB()
	users := make([]models.User, 0)

	var err error
	if email != "" {
		err = db.Find(&users, models.User{Email: email}).Error
	} else {
		err = db.Find(&users).Error
	}
	return users, err
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

func (UserService) AddUser(request validators.UserRegisterRequest) (models.User, error) {
	db := database.GetDB()

	user := validators.RegisterRequestToUserModel(request)

	err := db.Model(models.User{}).Create(&user).Error
	return user, err
}

func (UserService) UpdateUser(request validators.UserUpdateRequest, id int) (models.User, error) {
	db := database.GetDB()
	user := models.User{}

	if err := db.First(&user, id).Error; err != nil {
		return user, err
	}

	if request.Balance == nil && request.IsAdmin == nil {
		return user, nil
	}

	validators.UpdateRequestToUserModel(request, &user)

	if err := db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (UserService) DeleteUser(id uint) error {
	return nil
}
