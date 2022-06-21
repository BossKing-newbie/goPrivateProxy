package main

import (
	"flag"
	"go_private_proxy/constant"
	"go_private_proxy/routers"
	"log"
)

func main() {
	flag.Parse()
	log.Default().Println("go mod cache dir:", constant.GetYml().GetString("module.cache"))
	r := routers.InitializationRouter()
	_ = r.Run(":8138")
}
