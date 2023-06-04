package main

import (
	"documentation/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	router := setupRouter()

	// 建立一個新的 HTTP 測試器
	w := httptest.NewRecorder()

	requestBody := `{
		"EmployeeId": "40579",
		"Name": "John",
		"Email": "john@example.com",
		"Password": "password123",
		"Department": "IT"
	}`
	req, _ := http.NewRequest("POST", "/TSMC/users/register", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	defer w.Result().Body.Close()

	// 解析回應 JSON
	var response struct {
		Message string      `json:"message"`
		User    models.User `json:"User"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// 檢查回應內容是否符合預期
	assert.Equal(t, "Regist Successfully", response.Message)
	assert.Equal(t, "40579", response.User.EmployeeId)
	assert.Equal(t, "John", response.User.Name)
	assert.Equal(t, "john@example.com", response.User.Email)
	assert.Equal(t, "IT", response.User.Department)
}

func TestLoginUser(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()

	requestBody := `{
		"email": "john@example.com",
		"password": "password123"
	}`
	req, _ := http.NewRequest("POST", "/TSMC/users/login", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	defer w.Result().Body.Close()

	var response struct {
		Message  string      `json:"message"`
		User     models.User `json:"User"`
		Sessions string      `json:"Sessions"`
	}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Login Successfully", response.Message)
	assert.Equal(t, "john@example.com", response.User.Email)
	assert.NotEmpty(t, response.Sessions)
}

// func TestCreateDoc(t *testing.T) {
// 	router := setupRouter()

// 	// 模擬使用者註冊和登入
// 	simulateRegisterAndLogin(router)

// 	// 模擬新增文章請求
// 	createDocReq := httptest.NewRequest(http.MethodPost, "/TSMC/docs/createDoc", nil)
// 	createDocReq.Header.Set("Content-Type", "application/json")
// 	createDocReqBody := `{
// 		"title": "Sample Document",
// 		"content": "This is a sample document."
// 	}`
// 	createDocReq.Body = ioutil.NopCloser(strings.NewReader(createDocReqBody))

// 	createDocReq.Header.Set("Cookie", "40579")

// 	createDocRecorder := httptest.NewRecorder()
// 	router.ServeHTTP(createDocRecorder, createDocReq)
// 	assert.Equal(t, http.StatusOK, createDocRecorder.Code)
// 	defer createDocRecorder.Result().Body.Close()

// 	var response gin.H
// 	err := json.Unmarshal(createDocRecorder.Body.Bytes(), &response)
// 	if err != nil {
// 		t.Errorf("Failed to parse JSON response: %v", err)
// 	}

// 	expectedMessage := "Document created successfully"
// 	expectedTitle := "Test Document"
// 	expectedContent := "This is a test document"

// 	assert.Equal(t, expectedMessage, response["message"])
// 	assert.Contains(t, response, "newDoc")

// 	newDoc := response["newDoc"].(map[string]interface{})
// 	assert.Equal(t, expectedTitle, newDoc["title"])
// 	assert.Equal(t, expectedContent, newDoc["content"])
// }

// func simulateRegisterAndLogin(router *gin.Engine) string {
// 	// 模擬註冊請求
// 	registerReq := httptest.NewRequest(http.MethodPost, "/TSMC/users/register", nil)
// 	registerReq.Header.Set("Content-Type", "application/json")
// 	registerReqBody := `{
// 		"EmployeeId": "40579",
// 		"Name": "John",
// 		"Email": "john@example.com",
// 		"Password": "password123",
// 		"Department": "IT"
// 	}`
// 	registerReq.Body = ioutil.NopCloser(strings.NewReader(registerReqBody))
// 	registerRecorder := httptest.NewRecorder()
// 	router.ServeHTTP(registerRecorder, registerReq)
// 	defer registerRecorder.Result().Body.Close()

// 	// 模擬登入請求
// 	loginReqBody := `{
// 		"email": "john@example.com",
// 		"password": "password123"
// 	}`
// 	loginReq, _ := http.NewRequest("POST", "/TSMC/users/login", strings.NewReader(loginReqBody))
// 	loginReq.Header.Set("Content-Type", "application/json")

// 	loginRecorder := httptest.NewRecorder()
// 	router.ServeHTTP(loginRecorder, loginReq)
// 	defer loginRecorder.Result().Body.Close()

// 	var response struct {
// 		Message  string      `json:"message"`
// 		User     models.User `json:"User"`
// 		Sessions string      `json:"Sessions"`
// 	}
// 	json.Unmarshal(loginRecorder.Body.Bytes(), &response)

// 	return response.Sessions
// }
