package main

import (
	"fmt"
	"strconv"
	"time"
)

// 每隔一秒输出"Hello world"
func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test Hello world " + strconv.Itoa((i)))
		time.Sleep(time.Second)
	}

}

func main() {
	go test() // 开启了一个协程 主线程和协程同时执行

	for i := 0; i < 10; i++ {
		fmt.Println("main Hello golang " + strconv.Itoa((i)))
		time.Sleep(time.Second)
	}

}
