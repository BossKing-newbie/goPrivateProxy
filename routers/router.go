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
	webApi.GET("/module", api.ModuleApi)
	webApi.GET("/version", api.ModuleVersionApi)
	webApi.GET("/fileTree", api.CacheFileTreeApi)
	/*proxy模块*/
	r.Use(middleware.InitializationGoProxy())
	return r
}
