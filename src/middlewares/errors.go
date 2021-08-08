package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	BadRequest    = "BAD REQUEST"
	NotAuthorized = "NOT AUTHORIZED"
	NotFound       = "NOT FOUND"
	RecordNotFound = "record not found"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.ByType(gin.ErrorTypePublic).Last()
		if err != nil {
			log.Print(err.Err)

			switch err.Error() {
			case BadRequest:
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			case NotAuthorized:
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			case NotFound:
			case RecordNotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
			}
		}
	}
}
