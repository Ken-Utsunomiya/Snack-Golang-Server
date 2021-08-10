package handlers

import (
	"Snack-Golang-Server/src/models"
	"Snack-Golang-Server/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserList(c *gin.Context) {
	userService := services.UserService{}

	email := c.Query("email_address")
	userList, err := userService.GetUserList(email)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, gin.H{ "users": userList})
}

func UserRetrieve(c *gin.Context) {
	userService := services.UserService{}

	userId, _ := strconv.Atoi(c.Param("user_id"))
	user, err := userService.GetUser(userId)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, user)
}

func UserCommonList(c *gin.Context) {
	userService := services.UserService{}

	userCommonList, err := userService.GetUserCommonList()
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusOK, gin.H{ "users": userCommonList })
}

func UserCreate(c *gin.Context) {
	userService := services.UserService{}

	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	res, err := userService.AddUser(user)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func UserUpdate(c *gin.Context) {

}

func UserDelete(c *gin.Context) {

}
