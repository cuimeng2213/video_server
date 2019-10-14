package main
import(
	"net/http"
	"io"
	"github.com/julienschmidt/httprouter"
)
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	io.WriteString(w, "CreateUser Handler")

}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}


/*
handler ->validation{1.request, 2.user }->bussiness logist->response
1. data model 
2.error handle


session:功能实现
*/