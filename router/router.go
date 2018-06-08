package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nEdAy/wx_attendance_api_server/controller"
)

func init() {
	/*	ns := beego.NewNamespace("/v1",

			beego.NSNamespace("/user",
				beego.NSInclude(
					&controller.UserController{},
				),
			),
		)
		beego.AddNamespace(ns)
	*/

	// Disable Console Color
	// gin.DisableConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	//router := gin.Default()

	/*	userC := new(user.UserController)
		router.GET("/", getting)        // GetAll ...
		router.GET("/:id", getting)     // GetOne ...
		router.POST("/login/", posting) // Login 登录...
		// router.PUT("/", putting)        // Put ...
		router.DELETE("/:id", deleting) // Delete ...*/

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	//router.Run()
	// router.Run(":3000") for a hard coded port
}

/*
func InitRouter() *gin.Engine {
	router := gin.Default()
	// router.Static("/static", "/static")
	// router.StaticFS("/static", http.Dir("static"))
	// router.LoadHTMLGlob("templates/*")
	v1 := router.Group("/v1")
	// v0.Use(middlewares.Auth())
	{
		// curl -X POST http://127.0.0.1:8000/v0/member -d "login_name=hell31&password=g2223"
		v1.POST("/member", controller.MemberAdd)
		// curl -X GET http://127.0.0.1:8000/v0/member
		v1.GET("/member", controller.Register)
		// curl -X GET http://127.0.0.1:8000/v0/member/1
		v1.GET("/member/:id", controller.MemberGet)
		//curl -X PUT http://127.0.0.1:8000/v0/member/1 -d "login_name=haodaquan&password=1234"
		v1.PUT("/member/:id", controller.MemberEdit)
		// curl -X DELETE http://127.0.0.1:8000/v0/member/2
		v1.DELETE("/member/:id", controller.MemberDelete)
	}
	return router
}
*/
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1 := r.Group("/v1")
	{
		user := v1.Group("/user")
		{
			// curl -X POST https://127.0.0.1/v1/user/ -d ""
			user.POST("/", controller.Register)
			// curl -X POST https://127.0.0.1/v1/user/login/ -d ""
			user.POST("/login/", controller.Login)
			// curl -X GET  https://127.0.0.1/v1/user/
			user.GET("/", controller.UserList)
			// curl -X DELETE https://127.0.0.1/v1/user/login/1
			user.DELETE("/:id", controller.DelUser)
		}

	}

	return r
}
