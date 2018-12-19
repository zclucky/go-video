package session

import (
	"github.com/satori/go.uuid"
	"log"
	"sync"
	"time"
	"zc.com/video_server/api/dbops"
	"zc.com/video_server/api/defs"
)

type SimpleSession struct {
	Username string
	TTL      int64
}

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 100000
}

func deleteExpiredSession(sid string ) {
	sessionMap.Delete(sid)
	err := dbops.DeleteSession(sid)
	if err != nil {
		log.Printf("%s", err)
	}
}

func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(key, value interface{}) bool {
		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key, ss)
		return true
	})

}

func GenerateNewSessionId(un string) string {
	id, _ := uuid.NewV4()
	ttl := nowInMilli() + 30*60*1000 // 30分钟过期
	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)
	err := dbops.InsertSession(id.String(), ttl, un)
	if err != nil {
		return ""
	}
	return id.String()
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMilli()
		// 判断session 是否过期
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	}
	return "", true
}
