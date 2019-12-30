package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *http.Router {
	router := httprouter.New()
	router.GET("/videos/:vid-id", streamHandler)
	router.GET("/upload/:vid-id", uploadHandler)
	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":9090", r)
}
