package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"microservice/upload/util"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		log.Printf("File is too big")
		return
	}
	file, headers, err := r.FormFile("file")

	// 参数
	prj := r.FormValue("prj")
	tp := r.FormValue("tp")


	if prj != "" {
		// 建立项目目录
		dir := "./upload/video/"+r.FormValue("prj")
		res, _ := util.PathExists(dir)
		if !res {
			os.MkdirAll(dir, 777)
		}
	}

	if err != nil {
		log.Printf("Error when try to get file: %v", err)
		return
	}
	fmt.Println(headers.Header.Get("Content-Type"))
	//获取上传文件的类型
	if headers.Header.Get("Content-Type") != "image/png" {
		log.Printf("只允许上传png图片")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		return
	}
	fn := headers.Filename
	suffix := filepath.Ext(fn)
	// 获取当前日期毫秒
	FileName := util.Md5Crypt(strconv.FormatInt(time.Now().UnixNano(), 10), "HW", fn)
	FileName += suffix
	err = ioutil.WriteFile("./upload/video/"+FileName, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")

}
