package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goproxy/goproxy"
	"go_private_proxy/constant"
	"golang.org/x/mod/module"
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
		Cacher: goproxy.DirCacher(constant.GetYml().GetString("module.cache")),
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
		GetProxy().ServeHTTP(c.Writer, c.Request)
		name = strings.TrimPrefix(path.Clean(name), "/")
		fmt.Println("name:", name)
		getGoproxyCacheName(name)
	}
}

func getGoproxyCacheName(name string) {
	nameParts := strings.Split(name, "/@v/")
	if len(nameParts) != 2 {
		fmt.Println(nameParts)
	}

	if _, err := module.UnescapePath(nameParts[0]); err != nil {
		fmt.Println(err)
	}

	nameBase := path.Base(name)
	nameExt := path.Ext(nameBase)
	fmt.Println("nameBase:", nameBase)
	fmt.Println("nameExt:", nameExt)
	escapedModuleVersion := strings.TrimSuffix(nameBase, nameExt)
	moduleVersion, err := module.UnescapeVersion(escapedModuleVersion)
	fmt.Println("version", moduleVersion)
	if err != nil {
		fmt.Println(err)
	}

}
