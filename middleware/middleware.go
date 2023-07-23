package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neerubhandari/BlogPortal/utils"
)

func IsAuthenticate(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		c.Abort() // Prevents the request from being processed further
		return
	}

	_, err = utils.ParseJwt(cookie) // I'm assuming you want to parse the JWT token, not a string
	c.Next()
}
