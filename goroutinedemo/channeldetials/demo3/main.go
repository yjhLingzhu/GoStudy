package main

import (
	"fmt"
	"time"
)

// 函数1
func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world!", i)
	}
}

// 函数2
func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test 错误，", err)
		}
	}()
	// 定义一个map
	var myMap map[int]string
	myMap[0] = "yjh"
	// 没有make就使用了，会报错

}

// 如果一个协程出问题了，希望不要影响其他的协程，保证程序正常运行
// 使用defer + recover
func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("main: ", i)
	}
}
