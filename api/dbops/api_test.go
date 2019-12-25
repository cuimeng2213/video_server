package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func clearDB() {
	dbConn.Exec("truncate users")
}

func TestMain(m *testing.M) {
	clearDB()
	m.Run()
	//clearDB()
}

func TestApiFlow(t *testing.T) {
	t.Run("ADD", testAddUser)
	t.Run("Query", testQueryUser)
	t.Run("DEL", testDeleteUser)
	t.Run("REGET", testReQueryUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("tom", "12345")
	if err != nil {
		t.Errorf("%v\n", err.Error())
	}
}

func testQueryUser(t *testing.T) {
	pwd, err := GetUserCredential("tom")
	if err != nil {
		t.Errorf("%v\n", err.Error())
	}
	if pwd != "12345" {
		t.Errorf("%v\n", "pwd error")
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("tom", "12345")
	if err != nil {
		t.Errorf("Delete error%v\n", err)
	}
}

func testReQueryUser(t *testing.T) {
	pwd, err := GetUserCredential("tom")
	if err != nil {
		t.Errorf("%v\n", err.Error())
	}
	if pwd != "" {
		t.Errorf("%v\n", "pwd error")
	}
}

func TestComments(t *testing.T) {
	clearDB()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments", testAddComment)
	t.Run("ListComments", testListComments)
}

func testAddComment(t *testing.T) {
	vid := "123321"
	aid := 1
	text := "this is a test comments"
	err := AddNewComments(vid, aid, text)
	if err != nil {
		t.Errorf(">>>>%+v \n", err)
	}
}
func testListComments(t *testing.T) {
	vid := "123321"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/100000000, 10))
	res, err := ListComments(vid, from, to)
	fmt.Printf("AAAAAAAAAAAA: %d \n", len(res))
	if err != nil {
		t.Errorf(">>>>:%+v\n", err)
	}
	for i, ele := range res {
		fmt.Printf("###:%d %v \n", i, ele)
	}

}
