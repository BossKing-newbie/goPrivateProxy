package middleware

import (
	"github.com/goproxy/goproxy"
	"go_private_proxy/constant"
	"os"
)

func InitializationProxy() *goproxy.Goproxy {
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
