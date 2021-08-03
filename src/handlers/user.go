package handlers

import (
	"Snack-Golang-Server/src/models"
	"Snack-Golang-Server/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "UserList",
	})
}

func UserRetrieve(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "error": "UserID is invalid." })
		return
	}

	userService := services.UserService{}
	user, err := userService.GetUser(userId)
	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{ "error": "User not found." })
		return
	}
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "error": "Internal Server Error" })
		return
	}

	c.JSON(http.StatusOK, user)
}

func UserCommonList(c *gin.Context) {
	userService := services.UserService{}
	userCommonList, err := userService.GetUserCommonList()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "error": "Internal Server Error"})
	}

	c.JSON(http.StatusOK, gin.H{ "users": userCommonList })
}

func UserCreate(c *gin.Context) {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{ "error": "Bad Request"})
		return
	}

	userService := services.UserService{}
	res, err := userService.AddUser(user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{ "error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func UserUpdate(c *gin.Context) {

}

func UserDelete(c *gin.Context) {

}
