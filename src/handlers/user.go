package handlers

import (
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

}

func UserCreate(c *gin.Context) {

}

func UserUpdate(c *gin.Context) {

}

func UserDelete(c *gin.Context) {

}
