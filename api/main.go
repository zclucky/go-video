package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check session
	validateUserSession(r)

	m.r.ServeHTTP(w, r)
}

/**
	路由注册器，所有路由都在此注册
 */
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	/* 用户类api */
	// 创建用户
	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", Login)

	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	log.Println("server start at : 8000")
	log.Fatal(http.ListenAndServe(":8000", mh))
}
