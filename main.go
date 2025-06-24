package main

import (
	"go-gin-template/api"
	"go-gin-template/api/config"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Book API
// @version 1.0
// @description This is a sample book service API
// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
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
