package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var tempVid string

func clearTables() {
	dbConn.Exec("TRUNCATE users")
	dbConn.Exec("TRUNCATE videos")
	dbConn.Exec("TRUNCATE comments")
	dbConn.Exec("TRUNCATE sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWordFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredentail("zhangchao", "123")
	if err != nil {
		t.Errorf("Error of add user %v", err)
	}
}

func testGetUser(t *testing.T) {
	_, err := GetUserCredentail("zhangchao")
	if err != nil {
		t.Errorf("Error of get user %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("zhangchao", "123")
	if err != nil {
		t.Errorf("Error of delete user %v", err)
	}
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideo)
	t.Run("GetVideo", testGetVideo)
	t.Run("DelVideo", testDeleteVideo)
}

func testAddVideo(t *testing.T) {
	vi, err := AddVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideo : %v", err)
	}
	tempVid = vi.Id
}

func testGetVideo(t *testing.T) {
	_, err := GetVideo(tempVid)
	if err != nil {
		t.Errorf("Error of GetVideo : %v", err)
	}
}

func testDeleteVideo(t *testing.T) {
	err := DeleteVideo(tempVid)
	if err != nil {
		t.Errorf("Error of DeleteVideo : %v", err)
	}
}

func TestCommentWorkFlow(t *testing.T){
	clearTables()
	t.Run("PrepareUser",testAddUser)
	t.Run("PrepareVideo",testAddVideo)
	t.Run("AddComments",testAddNewComments)
	t.Run("ListComments",testListComments)

}

func testAddNewComments(t *testing.T) {
	err := AddNewComments(tempVid,1,"真好看")
	if err != nil {
		t.Errorf("Error of AddNewComments : %v", err)
	}
}
func testListComments(t *testing.T) {
	from := 1545153074
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := ListComments(tempVid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}
	for i, ele := range res {
		fmt.Printf("comment: %d, %v \n", i, ele)
	}
}





