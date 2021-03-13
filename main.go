package main

import (
	AuthController "Users/controllers/Auth"
	UserController "Users/controllers/User"
	"Users/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDataBase() // new

	r.GET("/users", UserController.FindUsers)
	r.GET("/users/:id", UserController.GetUser)
	r.PUT("/users/:id", UserController.UpdateUser)
	r.DELETE("/users/:id", UserController.DeleteUser)

	r.POST("/signup", AuthController.SignUp)
	r.POST("/signin", AuthController.SignIn)

	r.Run()
}
