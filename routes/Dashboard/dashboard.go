package Dashboard

import (
	"Users/middlewares/Auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	dashboard := route.Group("/dashboard")
	dashboard.Use(Auth.JwtVerify())
	{
		dashboard.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "OK"})
		})
	}
}
