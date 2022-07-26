package main


import (
	"fmt"
)

func test(n int)  {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n=", n)
}

// 函数递归，可以用栈的知识画图分析得出结果
func main()  {
	test(4)
}