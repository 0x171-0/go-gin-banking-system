```
mkdir go-gin-template
cd go-gin-template

# 初始化專案
go mod init go-gin-template

# 安裝 swagger 套件
go install github.com/swaggo/swag/cmd/swag@latest

go get -u github.com/swaggo/gin-swagger github.com/swaggo/files

# 安裝 jwt 套件
go get github.com/golang-jwt/jwt/v5

# 按裝測試套件
go get github.com/stretchr/testify
```
