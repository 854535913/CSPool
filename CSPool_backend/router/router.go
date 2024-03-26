package router

import (
	"github.com/gin-gonic/gin"
	"main/controller"
	mysqlmodule "main/dao/mysql"
	redismodule "main/dao/redis"
	"net/http"
)

func Init() *gin.Engine {
	r := gin.Default()
	dbMysql := mysqlmodule.Sdb
	dbRedis := redismodule.Rdb
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	userGroup := r.Group("/user")
	{
		userGroup.POST("/register", func(c *gin.Context) {
			controller.RegisterHandler(c, dbMysql)
		})
		userGroup.POST("/login", func(c *gin.Context) {
			controller.LoginHandler(c, dbMysql)
		})
		userGroup.GET("/info", controller.JWTAuthMiddleware(), func(c *gin.Context) {
			controller.GetUseriInfoHandler(c, dbMysql)
		})
	}

	postGroup := r.Group("/post", controller.JWTAuthMiddleware())
	{
		postGroup.POST("/upload", func(c *gin.Context) {
			controller.UploadHandler(c, dbMysql, dbRedis)
		})
		postGroup.GET("/:id", func(c *gin.Context) {
			controller.GetPostByIDHandler(c, dbMysql)
		})
		postGroup.GET("/time", func(c *gin.Context) {
			controller.GetPostlistByTimeHandler(c, dbMysql, dbRedis)
		})
		postGroup.GET("/like", func(c *gin.Context) {
			controller.GetPostlistByLikeHandler(c, dbMysql, dbRedis)
		})
		postGroup.GET("/review",
			func(c *gin.Context) {
				controller.CheckAdminAuthority(c, dbMysql)
			}, func(c *gin.Context) {
				controller.GetPostlistByUnderreviewHandler(c, dbMysql)
			})
		postGroup.POST("/publish/:id", func(c *gin.Context) {
			controller.ReviewHandler(c, dbMysql, dbRedis)
		})
		postGroup.POST("/:id/:action", func(c *gin.Context) {
			controller.VoteHandler(c, dbMysql, dbRedis)
		})
	}
	return r
}
