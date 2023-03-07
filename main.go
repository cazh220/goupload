package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goupload/routers"
	"log"
)

type Conf struct {
	host string
	port int
}
var conf Conf

func init()  {
	c, err := config.New(feeder.JsonDirectory{Path: "D:/Go/users/config"})
	if err != nil {
		log.Println(err)
	}

	conf.port, err = c.GetInt("app.port")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(conf.port)

}

func main() {
	// 创建路由
	r := gin.Default()
	// 注册路由
	routers.RegisterRouter(r)

	// 监听端口，默认在8080
	r.Run(":8000")
}
