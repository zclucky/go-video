package main

import (
	"net/http"
	"zc.com/video_server/api/defs"
	"zc.com/video_server/api/session"
)

var HEADER_FILED_SESSION = "X-Session-Id"
var HEADER_FILED_UNAME = "X-User-Name"


// 验证用户session
func validateUserSession(r *http.Request) bool{
	sid := r.Header.Get(HEADER_FILED_SESSION)
	if len(sid) == 0 {
		return false
	}

	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FILED_UNAME, uname)
	return true
}

// 验证用户
func ValidateUser(w http.ResponseWriter, r *http.Request) bool{
	uname := r.Header.Get(HEADER_FILED_UNAME)
	if len(uname) == 0 {
		SendErrorResponse(w,defs.ErrorNotAuthUser)
		return false
	}
	return true
}
