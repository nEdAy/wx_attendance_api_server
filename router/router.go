package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nEdAy/wx_attendance_api_server/controller"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/nEdAy/wx_attendance_api_server/docs"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())
	Router.Static("/assets", "./assets")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Ping test
	Router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	// v1
	v1 := Router.Group("/v1")
	{
		// 用户相关API
		user := v1.Group("/user")
		{
			// 注册用户 curl -X POST https://127.0.0.1/v1/user/ -d ""
			user.POST("/", controller.Register)
			// 用户登录 curl -X POST https://127.0.0.1/v1/user/login/ -d ""
			user.POST("/login/", controller.Login)
			// 获取全部用户 curl -X GET  https://127.0.0.1/v1/user/
			user.GET("/", controller.UserList)
			// 删除用户 curl -X DELETE https://127.0.0.1/v1/user/login/1
			user.DELETE("/:id", controller.DelUser)
		}
		// COS相关API
		cos := v1.Group("/cos")
		{
			// 获取鉴权签名 curl -X GET  https://127.0.0.1/v1/cos/
			cos.GET("/", controller.GetAuthorization)
		}
	}
}
