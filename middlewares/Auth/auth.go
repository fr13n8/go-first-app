package Auth

import (
	UserController "Users/controllers/User"
	"Users/httputil"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("Unathorized request!"))
			c.Abort()
			return
		}

		splitToken := strings.Split(token, "Bearer ")

		if len(splitToken) != 2 {
			httputil.NewError(c, http.StatusUnauthorized, errors.New("Broken auth token!"))
			c.Abort()
			return
		}
		token = splitToken[1]

		_, err := UserController.JwtVerify([]byte(token))

		if err != nil {
			httputil.NewError(c, http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		c.Next()
		return
	}
}
