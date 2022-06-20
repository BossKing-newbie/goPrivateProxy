package routers

import (
	"github.com/gin-gonic/gin"
	"go_private_proxy/middleware"
	"go_private_proxy/routers/api"
)

func InitializationRouter() *gin.Engine {
	r := gin.Default()
	/*proxy模块*/
	goProxy := r.Group("/")
	goProxy.Use(gin.WrapH(middleware.InitializationProxy()))
	/*api模块*/
	webApi := r.Group("/api")
	webApi.GET("/ping", api.TestApi)
	return r
}
