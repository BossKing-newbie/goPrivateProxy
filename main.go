package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goproxy/goproxy"
	"os"
)

func main() {
	r := gin.Default()
	r.GET("/", gin.WrapH(&goproxy.Goproxy{
		GoBinEnv: append(
			os.Environ(),
			"GOPROXY=https://goproxy.cn,direct", // 使用 Goproxy.cn 作为上游代理
		),
		ProxiedSUMDBs: []string{
			"sum.golang.org https://goproxy.cn/sumdb/sum.golang.org", // 代理默认的校验和数据库
		},
	}))
	_ = r.Run(":8080")
}
