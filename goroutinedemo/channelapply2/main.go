package main

import (
	"fmt"
	"math"
	"time"
)

// 计算1-8000之间哪些是素数

// 向numsChan管道写1-8000的数据
func writeNumToChan(numsChan chan int) {
	for i := 1; i <= 80000; i++ {
		numsChan <- i
	}
	// 写完数据关闭管道
	close(numsChan)
	fmt.Println("关闭numsChan")
}

// 计算素数
func prime(numsChan chan int, resultChan chan int, workChan chan bool) {
	// 获取数据
	for {
		v, ok := <-numsChan
		if !ok {
			// 完成工作了
			fmt.Println("完成工作啦！")
			workChan <- true
			break
		}
		if v == 1 {
			continue
		}
		if v == 2 {
			resultChan <- v // 将素数写到结果中
		} else {
			flag := true
			for i := 2; i <= int(math.Sqrt(float64(v))); i++ {
				if v%i == 0 {
					flag = false
					break
				}
			}
			if flag {
				resultChan <- v // 将素数写到结果中
			}
		}
	}
}

func main() {
	start := time.Now().Unix()
	// 新建一个管道，存放1-8000的数据
	numsChan := make(chan int, 1000)
	// 启动一个协程写入数据
	go writeNumToChan(numsChan)

	// 新建一个管道用于存放结果
	resultChan := make(chan int, 20000)
	// 新建一个管道用于记录四个协程是否完成工作
	workChan := make(chan bool, 4)

	//启4个计算素数的协程
	go prime(numsChan, resultChan, workChan)
	go prime(numsChan, resultChan, workChan)
	go prime(numsChan, resultChan, workChan)
	go prime(numsChan, resultChan, workChan)

	// 启一个匿名协程去监听啥时候关闭resultChan
	go func() {
		for {
			if len(workChan) == 4 {
				fmt.Println("关闭resultChan")
				end := time.Now().Unix()
				fmt.Println("总共耗时：", end-start)
				close(resultChan)
				break
			}
		}
	}()

	// 读取结果
	for {
		_, ok := <-resultChan
		if !ok {
			break
		}
		// fmt.Println(v)
	}
	fmt.Println("主线程退出")

}
