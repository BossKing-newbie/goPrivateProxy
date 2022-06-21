package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goproxy/goproxy"
	"go_private_proxy/constant"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
)

func GetProxy() *goproxy.Goproxy {
	/*继承了http.handler*/
	proxy := &goproxy.Goproxy{
		GoBinEnv: append(
			os.Environ(),
			"GOPROXY=https://goproxy.cn,direct", // 使用 Goproxy.cn 作为上游代理
		),
		ProxiedSUMDBs: []string{
			"sum.golang.org https://goproxy.cn/sumdb/sum.golang.org", // 代理默认的校验和数据库
		},
		Cacher: goproxy.DirCacher(*constant.ModuleCache),
	}
	return proxy
}

/*自定义中间件*/
func InitializationGoProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		name, err := url.PathUnescape(c.Request.URL.Path)
		if err != nil || strings.HasSuffix(name, "/") {
			c.Header(
				"Cache-Control",
				fmt.Sprintf("public, max-age=%d", 86400),
			)
			c.Status(http.StatusNotFound)
		}
		if path.Ext(name) != ".zip" {
			GetProxy().ServeHTTP(c.Writer, c.Request)
		}
	}
}
