package main

import (
	"fmt"
	"app"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	//统一异常处理
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("异常信息: %v \n", err)
		}
	}()
	app.StartUp()
}