package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckStatus(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
