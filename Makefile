# Go 相關變量
BINARY_NAME=go-gin-template
GO_FILES=$(wildcard *.go)

# Swagger 相關變量
SWAG_VERSION=latest
MIGRATE_VERSION=v4.17.0

# 數據庫連接信息
DB_URL=postgres://postgres:postgres@localhost:5432/bookstore?sslmode=disable

.PHONY: all build clean run swag-init swag-install help docker-build docker-run docker-stop docker-clean

# 顯示所有可用的命令
help:
	@echo "可用的命令："
	@echo "make build    - 編譯應用程序"
	@echo "make run      - 運行應用程序"
	@echo "make clean    - 清理編譯文件"
	@echo "make doc      - 生成 Swagger 文檔（需要先安裝 swag）"
	@echo "make install-swagger - 安裝 Swagger 工具"
	@echo "make install-migrate - 安裝 migrate 工具"
	@echo "make migrate-up     - 執行數據庫遷移"
	@echo "make migrate-down   - 回滾數據庫遷移"
	@echo "make migrate-create - 創建新的遷移文件"
	@echo "make all      - 清理、安裝依賴、生成文檔並編譯"
	@echo "make d-build  - 建立 Docker 映像"
	@echo "make up       - 運行 Docker 容器（包含 PostgreSQL 和 Redis）"
	@echo "make down     - 停止所有 Docker 容器"
	@echo "make clean    - 清理 Docker 容器和映像"

# 編譯應用程序
build:
	@echo "編譯應用程序..."
	go build -o $(BINARY_NAME) $(GO_FILES)

# 運行應用程序
run:
	@echo "運行應用程序..."
	go run main.go

# 清理編譯文件
go-clean:
	@echo "清理編譯文件..."
	go clean
	rm -f $(BINARY_NAME)
	rm -rf docs

# 安裝工具
install-swagger:
	@echo "安裝 Swagger 工具..."
	go install github.com/swaggo/swag/cmd/swag@$(SWAG_VERSION)

# 安裝 migrate 工具
install-migrate:
	@echo "安裝 migrate 工具..."
	# brew install golang-migrate
	# 或手動安裝
	# https://github.com/golang-migrate/migrate#installation
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@$(MIGRATE_VERSION)

# 數據庫遷移命令
migrate-up:
	@echo "執行數據庫遷移..."
	migrate -database "$(DB_URL)" -path scripts/migrations up

migrate-down:
	@echo "回滾數據庫遷移..."
	migrate -database "$(DB_URL)" -path scripts/migrations down

migrate-create: # 創建空的遷移文件，SQL 要自己填
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir scripts/migrations -seq $$name

# 生成 Swagger 文檔
doc:
	@echo "生成 Swagger 文檔..."
	swag init

# 執行所有步驟
all: clean install-swagger doc build

# Docker 相關命令
d-build:
	@echo "建立 Docker 映像..."
	docker-compose build

# 運行 Docker 容器
up:
	@echo "運行 Docker 容器..."
	docker-compose up -d
	@echo "服務啟動中，請稍候..."
	@sleep 5
	@echo "服務已啟動："
	@echo "- API: http://localhost:3003"
	@echo "- Swagger: http://localhost:3003/swagger/index.html"
	@echo "- PostgreSQL: localhost:5432"
	@echo "- Redis: localhost:6379"

# 停止 Docker 容器
down:
	@echo "停止 Docker 容器..."
	docker-compose down

# 清理 Docker 資源
clean:
	@echo "清理 Docker 資源..."
	docker-compose down -v --rmi all
	@echo "已清理所有 Docker 容器和映像"

# 如何使用 Swagger：
# 1. 首先執行 'make install-swagger' 安裝 Swagger 工具
# 2. 在程式碼中添加 Swagger 註解
# 3. 執行 'make swagger' 生成文檔
# 4. 運行應用程序後訪問 http://localhost:3002/swagger/index.html

# Swagger 註解示例：
# // @title Book API
# // @version 1.0
# // @description This is a sample book service API
# // @host localhost:3002
# // @BasePath /
#
# // @Summary Get all books
# // @Description Get a list of all books
# // @Tags books
# // @Accept json
# // @Produce json
# // @Success 200 {array} model.Book
# // @Router /books [get]
