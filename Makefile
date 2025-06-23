# Go 相關變量
BINARY_NAME=go-gin-template
GO_FILES=$(wildcard *.go)

# Swagger 相關變量
SWAG_VERSION=v1.16.4

.PHONY: all build clean run swag-init swag-install help

# 顯示所有可用的命令
help:
	@echo "可用的命令："
	@echo "make build    - 編譯應用程序"
	@echo "make run      - 運行應用程序"
	@echo "make clean    - 清理編譯文件"
	@echo "make swagger  - 生成 Swagger 文檔（需要先安裝 swag）"
	@echo "make install-swagger - 安裝 Swagger 工具"
	@echo "make all      - 清理、安裝依賴、生成文檔並編譯"

# 編譯應用程序
build:
	@echo "編譯應用程序..."
	go build -o $(BINARY_NAME) $(GO_FILES)

# 運行應用程序
run:
	@echo "運行應用程序..."
	go run main.go

# 清理編譯文件
clean:
	@echo "清理編譯文件..."
	go clean
	rm -f $(BINARY_NAME)
	rm -rf docs

# 安裝 Swagger 工具
install-swagger:
	@echo "安裝 Swagger 工具..."
	go install github.com/swaggo/swag/cmd/swag@$(SWAG_VERSION)

# 生成 Swagger 文檔
swagger:
	@echo "生成 Swagger 文檔..."
	swag init

# 執行所有步驟
all: clean install-swagger swagger build

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
