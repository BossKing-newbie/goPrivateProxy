package constant

import (
	"flag"
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var ModuleCache = flag.String("m", "/go/modules", "the module cache dir")

var (
	cMapSingleton sync.Once
	cMap          cmap.ConcurrentMap
	yml           *viper.Viper
	ymlSingleton  sync.Once
)

func initMap() cmap.ConcurrentMap {
	m := cmap.New()
	return m
}
func GetConcurrentMap() cmap.ConcurrentMap {
	cMapSingleton.Do(func() {
		cMap = initMap()
	})
	return cMap
}

// 读取yaml文件
func readYml() *viper.Viper {
	//读取yaml文件
	vB := viper.New()
	//设置读取的配置文件
	vB.SetConfigName("bootstrap")
	//设置配置文件类型
	vB.SetConfigType("yml")
	//添加读取的配置文件路径
	vB.AddConfigPath("./conf/")
	if err := vB.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
	}
	active := vB.GetString("active")
	log.Println("激活文件:", "config-"+active+".yml")
	//读取yaml文件
	vB = viper.New()
	//设置读取的配置文件
	vB.SetConfigName("config-" + active)
	//设置配置文件类型
	vB.SetConfigType("yml")
	//添加读取的配置文件路径
	vB.AddConfigPath("./conf/")
	if err := vB.ReadInConfig(); err != nil {
		log.Printf("err:%s\n", err)
	}
	return vB
}
func GetYml() *viper.Viper {
	ymlSingleton.Do(func() {
		yml = readYml()
	})
	return yml
}
