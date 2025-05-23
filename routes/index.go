package routes

import (
	"github.com/gin-gonic/gin"
	"loan-ext/app/http/api"
)

func IndexRoute(route *gin.RouterGroup) {
	IndexController := new(api.IndexController)
	route.GET("/", IndexController.Index)
}
