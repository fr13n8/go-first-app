package Messages

import (
	MessageController "Users/controllers/Messages"
	"Users/middlewares/Auth"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	message := route.Group("/message")
	message.Use(Auth.JwtVerify())
	{
		message.POST("/send", MessageController.SendMessage)
	}
}
