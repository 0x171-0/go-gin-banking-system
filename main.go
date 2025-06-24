package main

import (
	"go-gin-template/api"
	"go-gin-template/api/config"
	"go-gin-template/docs"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Gin API
// @version 1.0
// @description A RESTful API service with user authentication and account management
// @host localhost:3003
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// Initialize Swagger docs
	docs.SwaggerInfo.BasePath = "/"

	// Initialize database connection
	config.InitDB()

	// Initialize Redis connection
	config.InitRedis()

	// Initialize router
	r := api.InitRouter()

	// Swagger documentation endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Print Swagger documentation URL
	log.Printf("Swagger 文檔可在 http://localhost:3003/swagger/index.html 查看")

	if err := r.Run(":3003"); err != nil {
		log.Fatal(err)
	}
}
