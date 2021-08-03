package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	badRequest = "Bad Request"
	notAuthorized = "Not Authorized"
	notFound = "Not Found"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.ByType(gin.ErrorTypePublic).Last()
		if err != nil {
			log.Print(err.Err)

			switch err.Error() {
			case badRequest:
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			case notAuthorized:
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			case "Not Found":
			case "record not found":
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
