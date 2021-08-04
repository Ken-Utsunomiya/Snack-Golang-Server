package handlers

import (
	"Snack-Golang-Server/src/models"
	"Snack-Golang-Server/src/services"
	"fmt"
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
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	userService := services.UserService{}
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
	user := models.User{}
	err := c.Bind(&user)
	fmt.Println(user)
	if err != nil {
		c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	userService := services.UserService{}
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
