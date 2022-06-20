package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/goproxy/goproxy"
	"log"
	"os"
)

var moduleCache = flag.String("m", "/go/modules", "the module cache dir")

func main() {
	flag.Parse()
	log.Default().Println("go mod cache dir:", *moduleCache)
	/*继承了http.handler*/
	goproxyHandler := &goproxy.Goproxy{
		GoBinEnv: append(
			os.Environ(),
			"GOPROXY=https://goproxy.cn,direct", // 使用 Goproxy.cn 作为上游代理
		),
		ProxiedSUMDBs: []string{
			"sum.golang.org https://goproxy.cn/sumdb/sum.golang.org", // 代理默认的校验和数据库
		},
		Cacher: goproxy.DirCacher(*moduleCache),
	}
	r := gin.Default()
	goProxy := r.Group("/")
	goProxy.Use(gin.WrapH(goproxyHandler))
	ginWeb := r.Group("/api")
	ginWeb.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run(":8138")
}
