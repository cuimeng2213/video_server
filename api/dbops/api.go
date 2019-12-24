package dbops

import (
	"database/sql"
	"fmt"
	"goLangStudy/video_server/api/defs"
	"goLangStudy/video_server/api/utils"

	"time"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES(?,?)")

	if err != nil {
		fmt.Printf("Add user error : %v\n", err)
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
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

// 新增vido数据
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	// create uuid
	uuid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	fmt.Printf("%v \n", ctime)

	stmt, err := dbConn.Prepare(`INSERT INTO video_info(id, author_id,name, display_ctime) 
	VALUES(?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(uuid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	video_info := &defs.VideoInfo{
		Id:           uuid,
		AuthorId:     aid,
		Name:         name,
		DisplayCtime: ctime,
	}
	return video_info, nil
}

func GetVideo() (*defs.VideoInfo, error) {

}

func DeleteVideo() error {

}
