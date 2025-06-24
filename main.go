package main

import (
	"go-gin-template/api"
	"go-gin-template/api/config"
	_ "go-gin-template/docs"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Book API
// @version 1.0
// @description This is a sample book service API
// @host localhost:3003
// @BasePath /

func main() {
	// Initialize database connections
	config.InitDB()
	config.InitRedis()

	r := api.InitRouter()

	// Swagger documentation endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Print Swagger documentation URL
	log.Printf("Swagger 文檔可在 http://localhost:3003/swagger/index.html 查看")

	if err := r.Run(":3003"); err != nil {
		log.Fatal(err)
	}
}
