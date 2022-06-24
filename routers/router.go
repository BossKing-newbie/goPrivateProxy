package routers

import (
	"github.com/gin-gonic/gin"
	"go_private_proxy/middleware"
	"go_private_proxy/routers/api"
	"net/http"
)

func InitializationRouter() *gin.Engine {
	r := gin.Default()
	/*api模块*/
	webApi := r.Group("/api")
	webApi.GET("/module", api.ModuleApi)
	webApi.GET("/version", api.ModuleVersionApi)
	webApi.GET("/fileTree", api.CacheFileTreeApi)
	webApi.POST("/download", api.DownloadFile)
	webApi.GET("/dataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})
	/*proxy模块*/
	r.Use(middleware.InitializationGoProxy())
	return r
}
