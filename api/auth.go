package main

import (
	"goLangStudy/video_server/session"
	"net/http"
)

var HEADER_FILED_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

func validateUserSession(r *http.Request) bool {
	sid := r.header.Get(HEADER_FILED_SESSION)
	if len(sid) == 0 {
		return false
	}
	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Head.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		sendErrorResponse()
		return false
	}
	return true
}
