package dbops

import (
	"testing"
)

func clearDB() {
	dbConn.Exec("truncate users")
}

func TestMain(m *testing.M) {
	clearDB()
	m.Run()
	clearDB()
}

func TestApiFlow(t *testing.T) {
	t.Run("ADD", testAddUser)
	t.Run("Quert", testQueryUser)
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
