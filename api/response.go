package main

import (
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter) {
	io.WriteString(w, "")
}

func sendNormalRespnse(w http.ResponseWriter) {
	io.WriteString(w, "")
}
