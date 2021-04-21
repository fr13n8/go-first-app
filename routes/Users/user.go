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
		user.GET("/", UserController.GetUsers)
		user.GET("/:id", UserController.GetUser)
		user.PUT("/:id", UserController.UpdateUser)
		user.DELETE("/:id", UserController.DeleteUser)
	}
}
