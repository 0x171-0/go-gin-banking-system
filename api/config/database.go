package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"go-gin-template/api/model"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	// 嘗試按順序加載環境變量文件
	envFiles := []string{
		".env.local",    // 本地開發環境（不提交到版本控制）
		".env",          // 默認環境配置
		".env.example",  // 示例配置（如果沒有 .env）
	}

	var loaded bool
	for _, envFile := range envFiles {
		if err := godotenv.Load(envFile); err == nil {
			log.Printf("Loaded environment variables from %s", envFile)
			loaded = true
			break
		}
	}

	if !loaded {
		log.Println("No .env file found, using system environment variables")
	}
}

var (
	DB    *gorm.DB
	Redis *redis.Client
)

func InitDB() {
	// 檢查是否啟用自動遷移
	autoMigrate := getEnvOrDefault("AUTO_MIGRATE", "true") == "true"
	// 連接數據庫

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnvOrDefault("DB_HOST", "localhost"),
		getEnvOrDefault("DB_USER", "postgres"),
		getEnvOrDefault("DB_PASSWORD", "postgres"),
		getEnvOrDefault("DB_NAME", "bookstore"),
		getEnvOrDefault("DB_PORT", "5432"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db

	// 根據環境變量決定是否執行自動遷移
	if autoMigrate {
		log.Println("Automating database migration...")
		err = DB.AutoMigrate(
			&model.Book{},
			&model.Category{},
			&model.User{},
			&model.Order{},
			&model.OrderItem{},
			&model.Cart{},
			&model.CartItem{},
		)
		if err != nil {
			log.Fatalf("Failed to migrate database: %v", err)
		}
		log.Println("Database migration completed successfully")
	} else {
		log.Println("Skipping auto-migration (AUTO_MIGRATE=false)")
	}
}

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     getEnvOrDefault("REDIS_ADDR", "localhost:6379"),
		Password: getEnvOrDefault("REDIS_PASSWORD", ""),
		DB:       0,
	})
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
