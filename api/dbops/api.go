package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/satori/go.uuid"
	"log"
	"time"
	"zc.com/video_server/api/defs"
)

// 新增用户
func AddUserCredentail(loginName string, password string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users(login_name,password) values (?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, password)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

// 获取用户
func GetUserCredentail(loginName string) (string, error) {
	var (
		pwd string
		err error
	)

	stmtOut, err := dbConn.Prepare("SELECT password FROM users where login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

// 删除用户
func DeleteUser(loginName string, password string) error {
	stmtIns, err := dbConn.Prepare("DELETE FROM users where login_name = ? and password = ? ")
	if err != nil {
		log.Printf("Delete user error : %s", err)
		return err
	}
	_, err = stmtIns.Exec(loginName, password)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

// 添加视频
func AddVideo(aid int, name string) (*defs.Video, error) {
	vid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	createTtime := t.UnixNano()
	ctime := t.Format("Jan 02 2006, 15:04:05") // M D y , HH:MM:SS
	stmtIns, err := dbConn.Prepare(`INSERT INTO videos(id,author_id,name,display_create_time,create_time) VALUES(?,?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid.String(), aid, name, ctime, createTtime)
	if err != nil {
		return nil, err
	}
	res := &defs.Video{Id: vid.String(), AuthorId: aid, Name: name, DisplayCreateTime: ctime}
	defer stmtIns.Close()
	return res, nil
}

// 获取视频
func GetVideo(vid string) (*defs.Video, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_create_time FROM videos WHERE id=?")

	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &defs.Video{Id: vid, AuthorId: aid, Name: name, DisplayCreateTime: dct}
	return res, nil
}

// 删除视频
func DeleteVideo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM videos WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

func AddNewComments(vid string, aid int, content string) error {
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(` SELECT comments.id, users.Login_name, comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)

	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}

		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil
}
