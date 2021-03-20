package main

import (
	"Users/models"
	"Users/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	go models.ConnectDataBase()

	routes.InjectApi(r)

	r.Run()
}
