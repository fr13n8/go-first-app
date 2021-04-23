package Auth

import (
	AuthController "Users/controllers/Auth"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	auth := route.Group("/auth")
	{
		auth.POST("/signup", AuthController.SignUp)
		auth.POST("/signin", AuthController.SignIn)
		auth.GET("/verify/:token", AuthController.Verify)
	}
}
