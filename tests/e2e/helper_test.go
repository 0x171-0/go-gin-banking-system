package e2e

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

// testRequest 是一個輔助函數，用於發送測試請求
func testRequest(router http.Handler, method, path string, body interface{}) *httptest.ResponseRecorder {
	var reqBody *bytes.Buffer
	if body != nil {
		jsonBody, _ := json.Marshal(body)
		reqBody = bytes.NewBuffer(jsonBody)
	} else {
		reqBody = bytes.NewBuffer(nil)
	}

	req, _ := http.NewRequest(method, path, reqBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

// getAuthToken 是一個輔助函數，用於獲取測試用的認證 token
func getAuthToken(router http.Handler, email, password string) string {
	loginBody := map[string]interface{}{
		"email":    email,
		"password": password,
	}
	
	w := testRequest(router, "POST", "/api/v1/auth/login", loginBody)
	if w.Code != http.StatusOK {
		return ""
	}

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	token, ok := response["token"].(string)
	if !ok {
		return ""
	}
	return token
}
