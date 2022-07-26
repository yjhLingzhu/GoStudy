package main

import "fmt"

func main() {
	// 管道可以声明为只读或只写

	// 1. 默认情况下管道是双向的，可读可写
	// var chan1 chan int

	// 2. 声明为只写
	var chan2 chan<- int
	chan2 = make(chan<- int, 3)
	chan2 <- 3
	fmt.Println(len(chan2))

	// 3. 只读
	var chan3 <-chan int
	chan3 = make(<-chan int, 3) 
	fmt.Println(len(chan3))

}
