package session

import (
	"goLangStudy/video_server/api/defs"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionFromDB() {

}

func GenerateNewSession(un string) string {

}

func IsSessionExpired(sid string) (string, bool) {

}
