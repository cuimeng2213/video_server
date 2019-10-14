package main

import(
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()
	// 用户注册
	router.POST("/user", CreateUser)
	// 用户登录
	router.POST("/user:username", Login)
	return router
}


func main(){
	r := RegisterHandlers()
	http.ListenAndServe(":9090", r)
}