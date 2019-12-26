package dbops

import (
	"database/sql"
	"goLangStudy/video_server/api/defs"
	"strconv"
	//"sync"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlStr := strconv.FormatInt(ttl, 10)
	stmt, err := dbConn.Prepare("insert into sessions(session_id, TTL, login_name) values(?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sid, ttlStr, uname)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	session := &defs.SimpleSession{}

	stmt, err := dbConn.Prepare("select login_name, TTL from sessions where session_id=?")
	if err != nil {
		return nil, err
	}
	var ttl string
	var uname string
	err = stmt.QueryRow(sid).Scan(&uname, &ttl)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if ttlInt, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		session.TTL = ttlInt
		session.Username = uname
	} else {
		return nil, err
	}
	defer stmt.Close()
	return session, nil
}
