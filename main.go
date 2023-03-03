package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

const (
	MAX_UPLOAD_SIZE = 1024* 1024 * 200
)

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8080", r)
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/upload", UploadHandler)

	return router
}
