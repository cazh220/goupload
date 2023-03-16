package routers

import (
	"github.com/gin-gonic/gin"
	"goupload/controller"
)

func RegisterRouter(router *gin.Engine)  {
	RouterFile(router)
}

func RouterFile(engine *gin.Engine)  {
	var group = engine.Group("/file")
	{
		group.POST("/upload",controller.FileUpload)
		group.GET("/list", controller.ViewFiles)
		group.GET("/scan", controller.ScanDirs)
	}


}

