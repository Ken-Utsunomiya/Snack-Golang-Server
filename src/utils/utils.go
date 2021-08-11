package utils

import (
	"Snack-Golang-Server/src/models"
	"Snack-Golang-Server/src/validators"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func RequestToUserModel(request validators.UserRegisterRequest) models.User {
	user := models.User{}
	user.FirstName = request.FirstName
	user.LastName = request.LastName
	user.Email = request.Email
	user.ImageURI = request.ImageURI
	return user
}
