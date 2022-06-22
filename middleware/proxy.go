package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goproxy/goproxy"
	"go_private_proxy/constant"
	"golang.org/x/mod/module"
	"log"
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
		/*200意指download success计入统计范围*/
		if c.Writer.Status() == 200 {
			if validate, moduleV, mod := isStatistics(name); validate && len(mod) > 0 && len(moduleV) > 0 {
				statisticDownloadNum(mod + ":" + moduleV)
			}
		}
	}
}

/*是否属于版本下载行为*/
func isStatistics(name string) (validate bool, version string, mod string) {
	nameParts := strings.Split(name, "/@v/")
	if len(nameParts) != 2 {
		return false, "", ""
	}

	par, err := module.UnescapePath(nameParts[0])
	if err != nil {
		log.Fatal(err)
		return false, "", ""
	}
	/*获取后缀*/
	version = path.Base(name)
	if strings.Contains(version, "list") {
		return false, "", ""
	}
	if strings.Contains(version, ".zip") {
		return true, strings.ReplaceAll(version, ".zip", ""), par
	}
	return false, "", ""
}
func statisticDownloadNum(key string) {
	if constant.GetConcurrentMap().Has(key) {
		num, _ := constant.GetConcurrentMap().Get(key)
		count := num.(int64) + 1
		constant.GetConcurrentMap().Set(key, count)
	} else {
		var initCount int64 = 1
		constant.GetConcurrentMap().Set(key, initCount)
	}
}
