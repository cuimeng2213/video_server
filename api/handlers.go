package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "CreateUser Handler")

}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	fmt.Println(">>>>>>>>>>: ", uname)
	io.WriteString(w, uname)
}

/*
handler ->validation{1.request, 2.user }->bussiness logist->response
1. data model
2.error handle


session:功能实现
*/
