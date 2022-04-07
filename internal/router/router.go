package router

import (
	"account-module/internal/app/controller"
	"account-module/pkg/conf"
	"account-module/pkg/utils"
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

func NewRouter() *gin.Engine {
	g := gin.New()
	// 使用中间件
	//g.Use(middleware.NoCache)
	//g.Use(middleware.Options)
	//g.Use(middleware.Secure)
	//g.Use(middleware.Logging())
	//g.Use(middleware.RequestID())
	//g.Use(middleware.Prom(nil))
	//g.Use(middleware.Tracing("eagle-service"))
	//g.Use(mw.Translations())
	//
	//// load web router
	//LoadWebRouter(g)
	//
	//// 404 Handler.
	//g.NoRoute(api.RouteNotFound)
	//g.NoMethod(api.RouteNotFound)
	//
	//// 静态资源，主要是图片
	//g.Static("/static", "./static")
	//
	//// swagger api docs
	//g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//// pprof router 性能分析路由
	//// 默认关闭，开发环境下可以打开
	//// 访问方式: HOST/debug/pprof
	//// 通过 HOST/debug/pprof/profile 生成profile
	//// 查看分析图 go tool pprof -http=:5000 profile
	//// see: https://github.com/gin-contrib/pprof
	//// pprof.Register(g)
	//
	//// HealthCheck 健康检查路由
	//g.GET("/health", api.HealthCheck)
	//// metrics router 可以在 prometheus 中进行监控
	//// 通过 grafana 可视化查看 prometheus 的监控数据，使用插件6671查看
	//g.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// v1 router
	apiV1 := g.Group("/v1")
	apiV1.Use()
	{
		// 认证相关路由
		//apiV1.POST("/account", controller.CreateAccount)
		//apiV1.POST("/account/list", controller.CreateAccountList)
		//apiV1.POST("/v1/account/registered", controller.GetExistsAccounts)

		//apiV1.POST("/register", user.Register)
		//apiV1.POST("/login", user.Login)
		//apiV1.POST("/login/phone", user.PhoneLogin)
		//apiV1.GET("/vcode", user.VCode)
		//
		//// 用户
		//apiV1.GET("/users/:id", user.Get)
		//apiV1.Use(middleware.JWT())
		//{
		//	apiV1.PUT("/users/:id", user.Update)
		//	apiV1.POST("/users/follow", user.Follow)
		//	apiV1.GET("/users/:id/following", user.FollowList)
		//	apiV1.GET("/users/:id/followers", user.FollowerList)
		//}
	}

	return g
}

func Router() *gin.Engine {
	//if viper.GetString("app.mode") == "production" {
	//	gin.SetMode(gin.ReleaseMode)
	//	gin.DisableConsoleColor()
	//}

	var r = gin.Default()

	// 公共路由
	//PublicGroup := Router.Group("")
	//{
	//	router.InitTestRouter(PublicGroup)
	//}

	config := conf.Config
	if config.Application.LogLevel == "debug" {
		r.Use(utils.RequestLogger)
	}

	apiV1 := r.Group("/v1")
	apiV1.Use()

	{
		// 认证相关路由
		apiV1.POST("/account/create_one", controller.CreateAccount)
		apiV1.POST("/account/create_list", controller.CreateAccounts)
		apiV1.POST("/account/exist_list", controller.GetExistsAccounts)
		apiV1.POST("/account/find_one", controller.FindAccount)
		apiV1.POST("/account/find_list", controller.FindAccounts)
	}

	r.GET("/ping", func(c *gin.Context) {
		logger.Info("hello word")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
