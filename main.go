package main

import (
	"Users/docs"
	"Users/models"
	"Users/routes"
	"fmt"
	"log"
	"os"

	"Users/routes/Websocket"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @termsOfService http://swagger.io/terms/

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AppHost := os.Getenv("APP_HOST")
	AppPort := os.Getenv("APP_PORT")
	URL := fmt.Sprintf("%s:%s", AppHost, AppPort)

	docs.SwaggerInfo.Title = "Swagger api for chat"
	docs.SwaggerInfo.Description = "This is a sample server Chat server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s%s", os.Getenv("APP_PROTOCOL"), URL)
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()
	r.Use(cors.Default())

	go models.ConnectDataBase()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.InjectApi(r)

	go Websocket.H.Run()

	err = r.Run(URL)
	if err != nil {
		log.Fatal("Error to run: " + err.Error())
	}
}
