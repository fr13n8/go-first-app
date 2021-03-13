package routes

import (
	"Users/routes/Auth"
	"Users/routes/Dashboard"
	"Users/routes/Users"

	"github.com/gin-gonic/gin"
)

func InjectApi(route *gin.Engine) {
	Users.Routes(route)
	Auth.Routes(route)
	Dashboard.Routes(route)
}
