package dbops

import (
	"database/sql"
	"fmt"
	"goLangStudy/video_server/api/defs"
)

func AddUserCredential(loginName string, pwd string) error {
	smtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES(?,?)")

	if err != nil {
		fmt.Printf("Add user error : %v\n", err)
		return err
	}
	_, err = smtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer smtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	smtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		fmt.Printf("query pwd error: %v\n", err)
		return "", err
	}
	var pwd string
	err = smtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer smtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	smtDel, err := dbConn.Prepare("DELETE FROM users where login_name=? and pwd = ?")
	if err != nil {
		return err
	}
	_, err = smtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer smtDel.Close()
	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	// create uuid

	return nil, nil

}
