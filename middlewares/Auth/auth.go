package Auth

import (
	UserController "Users/controllers/User"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unathorized request"})
			return
		}

		splitToken := strings.Split(token, "Bearer ")
		token = splitToken[1]

		jh, err := UserController.JwtVerify([]byte(token))

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"data": jh})
		return
	}
}
