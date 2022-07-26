package main

import (
	"fmt"
)

// 可变参数
// args... 代表接受多个位置参数，类似python中的*args
func sum(n int, args... int) (res int) {
	res = n
	for i := 0; i < len(args); i++ {
		res += args[i]
		fmt.Println(i)
	}
	return
} 

// 交换两个数
func swap(n1, n2 int) (int, int) {
	n1, n2 = n2, n1
	return n1, n2
}

func main()  {
	res := sum(1,2,3,4,5)
	fmt.Println("sum=", res)

	res1, res2 := swap(2,3)
	fmt.Println(res1, res2)
}