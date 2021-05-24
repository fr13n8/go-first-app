package Websocket

import (
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	ws := route.Group("/ws")
	// ws.Use(Auth.JwtVerify())
	{
		ws.GET("/room/:id", func(c *gin.Context) {
			serveWs(c.Writer, c.Request, c.Param("id"))
		})
	}
}
