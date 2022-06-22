package api

import (
	"github.com/gin-gonic/gin"
	"go_private_proxy/service"
)

func ModuleVersionApi(c *gin.Context) {
	c.JSON(200, service.GetVersionList())
}
func ModuleApi(c *gin.Context) {
	c.JSON(200, service.GetModList())
}
