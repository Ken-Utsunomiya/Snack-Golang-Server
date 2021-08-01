package handlers

import (
	"Snack-Golang-Server/src/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "UserList",
	})
}

func UserRetrieve(c *gin.Context) {

}

func UserCommonList(c *gin.Context) {
	userService := services.UserService{}
	UserCommonList, err := userService.GetUserCommonList()
	if err != nil {
		c.JSON(http.StatusBadGateway, err)
	}
	c.JSON(http.StatusOK, UserCommonList)
}

func UserCreate(c *gin.Context) {

}

func UserUpdate(c *gin.Context) {

}

func UserDelete(c *gin.Context) {

}
