package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main()  {
	var count int = 0
	for {
		// 设置一个随机种子，不然的话是伪随机数
		rand.Seed(time.Now().UnixNano())
		// 随机数
		n := rand.Intn(100) + 1 
		// fmt.Println(n)
		count++
		if n == 99 {
			break
		}
	}
	fmt.Printf("生成99花费了%d次\n", count)
	

	// break可以跳出指定的层, break 默认跳出最近的for循环
	label2:
	for i := 1; i < 4; i++ {
		// label1:		// 为for循环添加label标签方便跳出
		for j := 0; j < 10; j++ {
			if j == 2 {
				// break		// 默认跳出
				// break label1		// 跳出label1的循环
				break label2  	// 跳出label2的循环
			}
			fmt.Println("j=", j)
		}
	}
	
}