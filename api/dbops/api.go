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

func GetVideo(vid string) (*defs.VideoInfo, error) {
	stmt, err := dbConn.Prepare("SELECT author_id, name, display_ctime from video_info where id=?")
	if err != nil {
		return nil, err
	}
	var aid int
	var name string
	var display_ctime string
	err = stmt.QueryRow(vid).Scan(&aid, &name, &display_ctime)
	if err != nil {
		return nil, err
	}

	return &defs.VideoInfo{
		Id:           vid,
		AuthorId:     aid,
		Name:         name,
		DisplayCtime: display_ctime,
	}, nil

}

// func DeleteVideo() error {

// }

func AddNewComments(vid string, aid int, text string) error {
	cid, err := utils.NewUUID()
	if err != nil {
		return err
	}
	stmt, err := dbConn.Prepare(`INSERT INTO comments(id,video_id,author_id,content) VALUES(?,?,?,?)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(cid, vid, aid, text)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	//
	/*stmtOut, err := dbConn.Prepare(`SELECT comments.id,users.login_name,comments.content from comments inner join
		users on comments.author_id = users.id
		where comments.video_id=? and comments.time > FROM_UNIXTIME(?) and comments.time <= FROM_UNIXTIME(?)
	`)*/
	stmtOut, err := dbConn.Prepare(`SELECT comments.id,users.login_name,comments.content from comments inner join
	users on comments.author_id = users.id where comments.video_id=?`)

	var res []*defs.Comment
	if err != nil {
		fmt.Printf("###########: %+v\n", err)
		return res, err
	}
	//rows, err := stmtOut.Query(vid, from, to)
	rows, err := stmtOut.Query(vid)
	if err != nil {
		return res, err
	}
	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		c := &defs.Comment{
			Id:      id,
			VideoId: vid,
			Author:  name,
			Content: content,
		}
		res = append(res, c)
	}
	defer stmtOut.Close()
	return res, nil
}
