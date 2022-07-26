package main

import (
	"fmt"
)


func main()  {

	// continue可以跳出指定的层, continue 默认跳出最近的for循环的本次
	label2:
	for i := 1; i < 4; i++ {
		// label1:		// 为for循环添加label标签方便跳出
		for j := 0; j < 10; j++ {
			if j == 2 {
				// break		// 默认跳出
				// break label1		// 跳出label1的本次循环
				continue label2  	// 跳出label2的本次循环
			}
			fmt.Println("j=", j)
		}
	}
	
}