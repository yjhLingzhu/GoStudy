package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 10
	intChan <- 30
	close(intChan) // 关闭
	// 此时这个channel只能读不能写
	num := <-intChan
	fmt.Println(num)

	// 管道的遍历只能使用for range，遍历的管道一定是要关闭的，对于没关闭的管道会报deathlock
	intChan2 := make(chan int, 100)
	for i := 1; i <= 100; i++ {
		intChan2 <- i * 2
	}
	close(intChan2)           // 关闭管道
	for v := range intChan2 { //管道没有下标的，所以只有一个v返回，平时的for range 是返回i,v的，所以这个决定返回是啥由range后面的对象决定
		fmt.Println(v)
	}
}
