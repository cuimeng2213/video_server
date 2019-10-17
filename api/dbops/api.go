package dbops

import (
	"fmt"
)

func AddUserCredential(loginName string, pwd string) error {
	smtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES(?,?)")

	if err != nil {
		fmt.Printf("Add user error : %v\n", err)
		return err
	}
	smtIns.Exec(loginName, pwd)
	smtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	smtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		fmt.Printf("query pwd error: %v\n", err)
		return "", err
	}
	var pwd string
	smtOut.QueryRow(loginName).Scan(&pwd)
	smtOut.Close()
	return pwd, nil
}

func DeletUser(loginName string, pwd string) error {
	smtDel, err := dbConn.Prepare("DELETE FROM users where login_name=? and pwd = ?")
	if err != nil {
		return err
	}
	smtDel.Exec(loginName, pwd)
	smtDel.Close()
	return nil
}
