package main

import (
	"fmt"
	"time"
)

// write data
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		// 写数据
		intChan <- i
		fmt.Println("写入数据：", i)
		time.Sleep(time.Second)
	}
	// 关闭
	close(intChan)

}

// read data
func readData(intChan chan int, exitChan chan bool) {
	// 读取
	// 读取数据时不受写数据的影响，即不管写数据的协程是否关闭管道
	// 都不受影响。在for循环中它会一直在那个管道中读取数据，只要管道关闭了，
	// 那么在取到最后值时就会返回false
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("读取数据：", v)
		time.Sleep(time.Second * 2)
	}
	// 读取完数据向exitChan写入数据
	exitChan <- true
	close(exitChan)
}

func main() {
	// 创建两个管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		// 这里相当于是阻塞，它会一直监听管道是否有数据，如果有就取出来返回去，
		// 否则就在那里等待，但是如果这个管道是关闭状态了的话，那么它将不会再
		// 等待下去。因为它认为不会有东西写进来了，所以由等到状态转成返回false
		// 的状态。
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
	// 读写的协程可以不同步，但是对于一个管道，一定要有读的协程，不然只有写的话它会报deadlock错误

}
