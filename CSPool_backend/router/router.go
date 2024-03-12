package router

import (
	"github.com/gin-gonic/gin"
	"main/controller"
	mysqlmodule "main/dao/mysql"
	"net/http"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.POST("/register", func(c *gin.Context) {
		controller.RegisterHandler(c, mysqlmodule.Sdb)
	})

	r.POST("/login", controller.LoginHandler)

	videoGroup := r.Group("/video", controller.JWTAuthMiddleware())
	{
		videoGroup.POST("/upload", controller.UploadHandler)
		videoGroup.GET("/:id", controller.GetVideoByIDHandler)
		videoGroup.GET("/time", controller.GetVideolistByTimeHandler)
		videoGroup.GET("/like", controller.GetVideolistByLikeHandler)
		videoGroup.POST("/review/:id", controller.ReviewHandler)
		videoGroup.POST("/:id/:action", controller.VoteHandler)
	}
	return r
}
