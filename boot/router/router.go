package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/small-ek/antgo/frame/gin_middleware"
	"github.com/small-ek/antgo/os/config"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"io/ioutil"
	"loan-ext/app/http/middleware"
	_ "loan-ext/docs"
	"loan-ext/routes"
)

func Router() *gin.Engine {
	var app = gin.New()
	//开发者模式
	if config.GetBool("system.debug") == false {
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	} else {
		app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	app.Use(requestid.New()).Use(gin_middleware.Recovery()).Use(gin_middleware.Logger())
	//跨域处理
	if config.GetBool("system.cors") == true {
		corsConfig := cors.DefaultConfig()
		corsConfig.AllowOrigins = []string{"*"}
		corsConfig.AllowHeaders = []string{"*"}
		app.Use(cors.New(corsConfig))
	}
	return app
}

// Load 加载路由
func Load() *gin.Engine {
	app := Router()
	//添加路由组前缀
	Group := app.Group("api")
	//注册路由
	routes.IndexRoute(Group)

	Group.Use(middleware.AuthJWT())
	{

	}

	return app
}
