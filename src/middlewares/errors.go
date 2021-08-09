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
	InternalServerError = "Internal Server Error"
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
					"error": BadRequest,
				})
			case NotAuthorized:
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"error": NotAuthorized,
				})
			case NotFound:
			case RecordNotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"error": NotFound,
				})
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": InternalServerError,
				})
			}
		}
	}
}
