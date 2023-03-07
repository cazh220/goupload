package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goupload/routers"
	"gopkg.in/ini.v1"
	"os"
)

type Conf struct {
	Host   string
	Port   string
	DbHost string
	DbPort string
	DbDriver string
	DbName string
}
var conf Conf

func init()  {
	fmt.Println(456)
	cfg, err := ini.Load("./config/my.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	conf.DbDriver = cfg.Section("db").Key("driver").String()
	conf.DbHost = cfg.Section("db").Key("host").String()
	conf.DbName = cfg.Section("db").Key("name").String()
	conf.DbPort = cfg.Section("db").Key("port").String()
	conf.Host = cfg.Section("").Key("host").String()
	conf.Port = cfg.Section("").Key("port").String()
}

func main() {
	// 创建路由
	r := gin.Default()
	// 注册路由
	routers.RegisterRouter(r)
	// 监听端口，默认在8080
	r.Run(":"+conf.Port)
}
