package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

var a int

func main(){
	fmt.Println(time.Now(),reflect.TypeOf(time.Now()))
	fmt.Println(time.Now().UnixNano()/1000000000,reflect.TypeOf(time.Now().UnixNano()/1000000000))
	fmt.Println(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10),reflect.TypeOf(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10)))
	a,_ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	fmt.Println(a,reflect.TypeOf(a))

	fmt.Println(runtime.Compiler, runtime.GOARCH, runtime.GOOS)
	fmt.Println(strconv.IntSize)
}

func test1(){
	a = 10
}

func test2() int {
	return a
}
