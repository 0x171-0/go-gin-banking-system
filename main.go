package main

import (
	"context"
	"go-gin-template/api"
	"go-gin-template/api/config"
	"go-gin-template/api/service"
	"go-gin-template/docs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

// 全局变量，用于在应用程序关闭时访问通知服务
var notificationService service.NotificationService

func main() {
	// Initialize Swagger docs
	docs.SwaggerInfo.BasePath = "/"

	// Initialize database connection
	config.InitDB()

	// Initialize Redis connection
	config.InitRedis()

	// 初始化通知服务
	notificationService = service.NewNotificationService()

	// Initialize router
	r := api.InitRouter()

	// Swagger documentation endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Print Swagger documentation URL
	log.Printf("Swagger 文檔可在 http://localhost:3003/swagger/index.html 查看")

	// 创建 HTTP 服务器
	srv := &http.Server{
		Addr:    ":3003",
		Handler: r,
	}

	// 在单独的 goroutine 中启动服务器
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("启动服务器失败: %v", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	// kill (无参数) 默认发送 syscall.SIGTERM
	// kill -2 是 syscall.SIGINT
	// kill -9 是 syscall.SIGKILL 但无法被捕获，所以不需要添加
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("正在关闭服务器...")

	// 设置 5 秒的超时时间来关闭服务器
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 等待所有邮件发送完成
	log.Println("等待所有邮件发送完成...")
	notificationService.WaitForCompletion()
	log.Println("所有邮件已发送完成")

	// 关闭 HTTP 服务器
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("服务器强制关闭:", err)
	}

	log.Println("服务器已优雅关闭")
}
