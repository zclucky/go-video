package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

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
	log.Println("server start at : 8000")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		panic("服务启动失败!")
	}
}

