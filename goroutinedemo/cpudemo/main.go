package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// demo1()
	demo2()
}

func demo1() {
	// 获取cpu个数
	cpuNum := runtime.NumCPU()
	fmt.Println("cpu个数：", cpuNum)

	// 设置cpu个数
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}

// 编写一个函数，计算各个数的阶乘，并放到map中
// 启动多个协程，统计的结果放到map中
// map应该是全局的

var (
	myMap = make(map[int]int, 10)
	// lock 是全局的互斥锁
	lock sync.Mutex
)

func demo2() {
	// 开启多个协程
	for i := 1; i <= 10; i++ {
		go cal(i)
	}

	// 休眠10秒钟
	time.Sleep(time.Second * 10)

	// 这里输出结果, 这里读的时候也加锁了
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%v]=%v\n", i, v)
	}
	lock.Unlock()
}

// 计算阶乘函数
func cal(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	// 加锁
	lock.Lock()
	myMap[n] = res // 将结果存放到map中
	// 解锁
	lock.Unlock()
}

// 加锁只能保证资源只有一个协程在用，并不能保证协程执行的顺序
