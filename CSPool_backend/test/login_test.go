package test

import (
	"bytes"
	"encoding/json"
	"main/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	if err := InitMySQL(); err != nil {
		t.Fatalf("MySQL start failed, err:%v\n", err)
	}
	defer sdbTest.Close()
	url := "/login"
	r.POST(url, func(c *gin.Context) {
		controller.LoginHandler(c, sdbTest)
	})

	body := `{
		"username" : "xiaohao",
        "password" : "1234"
	}`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	res := new(controller.ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, controller.CodeSuccess, res.Code)
}

func TestLoginHandler_UserNotExist(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	if err := InitMySQL(); err != nil {
		t.Fatalf("MySQL start failed, err:%v\n", err)
	}
	defer sdbTest.Close()
	url := "/login"
	r.POST(url, func(c *gin.Context) {
		controller.LoginHandler(c, sdbTest)
	})

	body := `{
		"username" : "NotExistUser",
        "password" : "1234"
	}`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	res := new(controller.ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, controller.CodeUserNotExist, res.Code)
}

func TestLoginHandler_InvalidPassword(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	if err := InitMySQL(); err != nil {
		t.Fatalf("MySQL start failed, err:%v\n", err)
	}
	defer sdbTest.Close()
	url := "/login"
	r.POST(url, func(c *gin.Context) {
		controller.LoginHandler(c, sdbTest)
	})

	body := `{
		"username" : "xiaohao",
        "password" : "InvalidPassword"
	}`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	res := new(controller.ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, controller.CodeInvalidPassword, res.Code)
}
