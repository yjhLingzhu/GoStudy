package main

import "fmt"

func main() {
	// 使用select可以解决从管道中获取数据的阻塞问题

	// 定义一个管道 10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	// 定义一个管道 5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "yjh_" + fmt.Sprintf("%d", i)
	}

	// 传统的方法在遍历管道时，如果不关闭会阻塞而导致 deadlock
	// 在实际开发中我们不确定什么时候关闭管道
	// 所以可以使用select方法解决
label1:
	for {
		select {
		// 注意：这里如果intChan一直没有关闭，不会一直阻塞而报deadlock，
		// 它会自动执行下一个case
		case v := <-intChan:
			fmt.Println("在intChan中取到了数据：", v)
		case v := <-stringChan:
			fmt.Println("在stringChan中取到了数据：", v)
		default:
			fmt.Println("都没取到数据，程序员可以在这里处理自己的逻辑")
			break label1
			// 这里也可也用return，这样的话是直接退出整个函数，单单使用break是跳不出for的
		}
	}
}
