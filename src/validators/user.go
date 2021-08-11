package validators

import "Snack-Golang-Server/src/models"

type UserRegisterRequest struct {
	FirstName	string `json:"first_name" binding:"required"`
	LastName	string `json:"last_name" binding:"required"`
	Email			string `json:"email_address" binding:"required"`
	ImageURI	string `json:"image_uri" binding:"required"`
}

func RegisterRequestToUserModel(request UserRegisterRequest) models.User {
	user := models.User{}
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email
	user.ImageURI = request.ImageURI
	return user
}
