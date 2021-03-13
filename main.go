package main

import (
	"Users/models"
	"Users/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDataBase() // new

	routes.InjectApi(r)

	r.Run()
}
