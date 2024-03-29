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

func TestRegisterHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	if err := InitMySQL(); err != nil {
		t.Fatalf("MySQL start failed, err:%v\n", err)
	}
	defer sdbTest.Close()
	url := "/register"
	r.POST(url, func(c *gin.Context) {
		controller.RegisterHandler(c, sdbTest)
	})

	body := `{
		"username" : "user3",
        "password" : "1234",
        "re_password":"1234"
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

func TestRegisterHandler_UsernameExist(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	if err := InitMySQL(); err != nil {
		t.Fatalf("MySQL start failed, err:%v\n", err)
	}
	defer sdbTest.Close()
	url := "/register"
	r.POST(url, func(c *gin.Context) {
		controller.RegisterHandler(c, sdbTest)
	})

	body := `{
		"username" : "user1",
        "password" : "1234",
        "re_password":"1234"
	}`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	res := new(controller.ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, controller.CodeUserExist, res.Code)
}
