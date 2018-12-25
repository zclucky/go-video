package main

import (
	"fmt"
	"zc.com/video_server/test/interface/mock"
)

// 定义接口
type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	var r Retriever
	r = &mock.Retriever{"this is new comments"}
	fmt.Printf("%T + %v", r, r)
	fmt.Println()
	fmt.Println(download(r))
}
