package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"

	"go-gin-template/api"
	"go-gin-template/api/config"
)

type AuthTestSuite struct {
	suite.Suite
	router http.Handler
}

func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}

func (s *AuthTestSuite) SetupSuite() {
	// 載入測試環境配置
	err := godotenv.Load("../../tests/e2e/.env.test")
	if err != nil {
		s.T().Logf("Warning: .env.test file not found: %v", err)
	}

	// 初始化資料庫連接
	config.InitDB()
	config.InitRedis()
	
	// 初始化路由
	router := api.InitRouter()
	s.router = router
}

func (s *AuthTestSuite) SetupTest() {
	// 在每個測試前清理數據
	s.cleanTestData()
}

func (s *AuthTestSuite) TearDownTest() {
	// 在每個測試後清理數據
	s.cleanTestData()
}

func (s *AuthTestSuite) TearDownSuite() {
	// 在所有測試完成後清理數據並關閉連接
	s.cleanTestData()
	
	// 關閉數據庫連接
	if config.DB != nil {
		if sqlDB, err := config.DB.DB(); err == nil {
			sqlDB.Close()
		}
	}
	
	// 關閉 Redis 連接
	if config.Redis != nil {
		config.Redis.Close()
	}
}

func (s *AuthTestSuite) cleanTestData() {
	db := config.DB
	// 按順序刪除相關的外鍵數據
	db.Exec("DELETE FROM accounts WHERE user_id IN (SELECT id FROM users WHERE email LIKE 'test%@example.com')")
	db.Exec("DELETE FROM user_passwords WHERE user_id IN (SELECT id FROM users WHERE email LIKE 'test%@example.com')")
	db.Exec("DELETE FROM users WHERE email LIKE 'test%@example.com'")
}

func (s *AuthTestSuite) TestRegisterAndLogin() {
	// 測試註冊
	registerURL := "/users/register"
	registerBody := map[string]interface{}{
		"email":     "test@example.com",
		"password":  "Test123!@#",
		"name":      "Test User",
	}
	jsonBody, _ := json.Marshal(registerBody)
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", registerURL, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusCreated, w.Code)
	
	// 測試登入
	loginURL := "/users/login"
	loginBody := map[string]interface{}{
		"email":    "test@example.com",
		"password": "Test123!@#",
	}
	jsonBody, _ = json.Marshal(loginBody)
	
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", loginURL, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusOK, w.Code)
	
	// 解析回應
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	s.NoError(err)
	
	// 確認回應中包含 token
	token, exists := response["token"]
	s.True(exists)
	s.NotEmpty(token)
}

func (s *AuthTestSuite) TestInvalidLogin() {
	loginURL := "/users/login"
	loginBody := map[string]interface{}{
		"email":    "nonexistent@example.com",
		"password": "wrongpassword",
	}
	jsonBody, _ := json.Marshal(loginBody)
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", loginURL, bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	s.router.ServeHTTP(w, req)

	s.Equal(http.StatusUnauthorized, w.Code)
}
