package Users

import (
	UserController "Users/controllers/User"
	"Users/middlewares/Auth"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	user := route.Group("/users")
	user.Use(Auth.JwtVerify())
	{
		user.GET("/", UserController.GetUsers).Use(Auth.JwtVerify())
		user.GET("/:id", UserController.GetUser).Use(Auth.JwtVerify())
		user.PUT("/:id", UserController.UpdateUser).Use(Auth.JwtVerify())
		user.DELETE("/:id", UserController.DeleteUser).Use(Auth.JwtVerify())
	}
}
