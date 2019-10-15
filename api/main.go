package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	// 用户注册
	router.POST("/user", CreateUser)
	// 用户登录 -- 使用path参数
	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":9090", r)
}
