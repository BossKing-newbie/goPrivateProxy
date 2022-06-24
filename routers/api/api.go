package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_private_proxy/constant"
	"go_private_proxy/service"
	"log"
	"net/http"
	"os"
)

func ModuleVersionApi(c *gin.Context) {
	c.JSON(200, service.GetVersionList())
}
func ModuleApi(c *gin.Context) {
	c.JSON(200, service.GetModList())
}
func CacheFileTreeApi(c *gin.Context) {
	c.JSON(200, service.ListFileCache(constant.GetYml().GetString("module.cache"), 1))
}
func DownloadFile(c *gin.Context) {
	var downloadModel service.DownloadModel
	e := c.ShouldBindJSON(&downloadModel)
	if e != nil {
		log.Fatal("ctx.ShouldBindJSON err: ", e)
		return
	}
	//打开文件
	file, errByOpenFile := os.Open(downloadModel.Path)
	//非空处理
	if errByOpenFile != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "失败",
			"error":   "资源不存在",
		})
	}
	reader := file
	fi, _ := file.Stat()
	contentLength := fi.Size()
	contentType := "application/octet-stream"
	fileVal := fmt.Sprintf("attachment; filename=%s", file.Name())
	extraHeaders := map[string]string{
		"Content-Disposition": fileVal,
	}

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}
