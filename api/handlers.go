package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"zc.com/video_server/api/dbops"
	"zc.com/video_server/api/defs"
	"zc.com/video_server/api/session"
)

// 新建用户
func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res, ubody); err != nil {
		SendErrorResponse(w, defs.ErrorRequestBodyParseFailed) // 参数错误
		return
	}

	if err := dbops.AddUserCredentail(ubody.Username, ubody.Password); err != nil {
		SendErrorResponse(w, defs.ErrorDbError) // 数据库错误
		return
	}

	// 生成session
	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success: true, SessionId: id}

	if resp, err := json.Marshal(su); err != nil {
		SendErrorResponse(w, defs.ErrorInternalFaults) // 网络错误
		return
	} else {
		SendNormalResponse(w, string(resp), 201)
	}
}

// 用户登录
func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	_, err := io.WriteString(w, uname)
	if err != nil {
		panic("登录出错")
	}
}
