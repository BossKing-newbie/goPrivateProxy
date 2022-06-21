package routers

import (
	"github.com/gin-gonic/gin"
	"go_private_proxy/middleware"
	"go_private_proxy/routers/api"
)

func InitializationRouter() *gin.Engine {
	r := gin.Default()
	/*api模块*/
	webApi := r.Group("/api")
	webApi.GET("/ping", api.TestApi)
	/*proxy模块*/
	r.Use(middleware.InitializationGoProxy())
	return r
}
