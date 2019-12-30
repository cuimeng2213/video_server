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
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		fmt.Printf(">>>Error:%v\n", err)
		return
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k, ss)
		return true
	})
}
func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}
func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dops.DeleteSession(sid)
}
func GenerateNewSession(un string) string {
	id, _ := utils.NewUUID()
	ct := nowInMilli()
	ttl := ct + 30*60*1000 //超时时间为30分钟
	ss := &defs.SimpleSession{
		Username: un,
		TTL:      ttl,
	}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)
	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	}
	return "", true
}
