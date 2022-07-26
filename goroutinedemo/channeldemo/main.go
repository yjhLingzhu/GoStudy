package main

import "fmt"

func main() {
	// 创建一个管道
	var intChan chan int
	intChan = make(chan int, 3)

	// 向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num

	// 看看管道的长度和cap(容量)
	fmt.Println(len(intChan), cap(intChan))

	var num2 int
	num2 = <-intChan // 在没有使用协程的情况下，如果我们的管道数据已经
	// 全部取出，再取就会报错 deadlock
	fmt.Println(num2)
}
