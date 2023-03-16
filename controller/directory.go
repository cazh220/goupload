package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func ScanDirs(context *gin.Context) {
	pwd, _ := os.Getwd()

	//filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
	//	fmt.Println(path)
	//	fmt.Println(info.Name())
	//	return nil
	//})
	fileInfoList, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(fileInfoList))

	for i := range fileInfoList {
		fmt.Println(fileInfoList[i].Name())
	}


	context.JSON(http.StatusOK, gin.H{
		"code": 10000,
		"msg":  "扫描成功",
		"data": nil,
	})

	return
}