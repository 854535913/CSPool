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

func TestRegisterHandler(t *testing.T) {
	if err := MySQLInit(); err != nil {
		t.Fatalf("MySQL start failed, err:%v\n", err)
	}
	// 确保在测试结束时关闭数据库连接
	defer Sdb_test.Close()

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/register"
	r.POST(url, func(c *gin.Context) {
		controller.RegisterHandler(c, Sdb_test)
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

	// 判断响应的内容是不是按预期返回了需要登录的错误
	res := new(controller.ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, controller.CodeSuccess, res.Code)
}
