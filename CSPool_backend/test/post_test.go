package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"main/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func LoginAndGetToken(r *gin.Engine, isadmin bool) string {
	url := "/login"
	r.POST(url, func(c *gin.Context) {
		controller.LoginHandler(c, sdbTest)
	})
	var body string
	if isadmin {
		body = `{
		"username" : "admin",
        "password" : "1234"
	}`
	} else {
		body = `{
		"username" : "user",
        "password" : "1234"
	}`
	}
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	res := new(controller.ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		fmt.Printf("解析响应体失败: %v\n", err)
		return ""
	}

	// 由于data是interface{}类型，我们需要根据实际类型进行断言
	var token string
	switch v := res.Data.(type) {
	case string:
		token = v
	default:
		fmt.Println("Token数据格式不是预期的string类型")
		return ""
	}
	return "Bearer " + token
}

func TestUploadHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	if err := InitMySQL(); err != nil {
		t.Fatalf("MySQL start failed, err:%v\n", err)
	}
	defer sdbTest.Close()
	if err := InitRedis(); err != nil {
		t.Fatalf("Redis start failed, err:%v\n", err)
	}
	defer rdbTest.Close()

	url := "/post/upload"
	r.POST(url, controller.JWTAuthMiddleware(), func(c *gin.Context) {
		controller.UploadHandler(c, sdbTest, rdbTest)
	})

	body := `{
    "title" : "test",
    "description" : "it is a test",
    "link": "test.com",
    "tag" : "test"
}`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	req.Header.Set("Authorization", LoginAndGetToken(r, false))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	res := new(controller.ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, controller.CodeSuccess, res.Code)
}

func TestUploadHandler_NeedLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	if err := InitMySQL(); err != nil {
		t.Fatalf("MySQL start failed, err:%v\n", err)
	}
	defer sdbTest.Close()
	if err := InitRedis(); err != nil {
		t.Fatalf("Redis start failed, err:%v\n", err)
	}
	defer rdbTest.Close()

	url := "/post/upload"
	r.POST(url, controller.JWTAuthMiddleware(), func(c *gin.Context) {
		controller.UploadHandler(c, sdbTest, rdbTest)
	})

	body := `{
    "title" : "bing",
    "description" : "bing official website",
    "link": "bing.com",
    "tag" : "website"
}`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	res := new(controller.ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, controller.CodeNeedLogin, res.Code)
}

func TestUploadHandler_TitleExist(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	if err := InitMySQL(); err != nil {
		t.Fatalf("MySQL start failed, err:%v\n", err)
	}
	defer sdbTest.Close()
	if err := InitRedis(); err != nil {
		t.Fatalf("Redis start failed, err:%v\n", err)
	}
	defer rdbTest.Close()

	url := "/post/upload"
	r.POST(url, controller.JWTAuthMiddleware(), func(c *gin.Context) {
		controller.UploadHandler(c, sdbTest, rdbTest)
	})

	body := `{
    "title" : "bing",
    "description" : "exist title",
    "link": "bing.com",
    "tag" : "website"
}`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	req.Header.Set("Authorization", LoginAndGetToken(r, false))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	res := new(controller.ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, controller.CodeTitleExist, res.Code)
}
