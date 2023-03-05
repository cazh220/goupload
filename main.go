package main

import (
	"github.com/gin-gonic/gin"
	"goupload/routers"
)

func main() {
	// 创建路由
	r := gin.Default()
	// 注册路由
	routers.RegisterRouter(r)

	// 监听端口，默认在8080
	r.Run(":8000")
}
