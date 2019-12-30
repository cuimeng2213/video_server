package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandle(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}
func (m middleWareHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	//检查sessionid合法性
	m.r.ServerHTTP(w, r)
}

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
	mh := NewMiddleWareHandle(r)
	http.ListenAndServe(":9090", mh)
}
