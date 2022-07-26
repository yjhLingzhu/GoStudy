package main

import (
	"fmt"
)

// main函数是这个go的函数入口，它一定要在main包里面
func main()  {
	fmt.Println(5/2)
	// 取模的公式：a % b = a - a/b * b
	// i++, i--是独立的，不能 a = i++ 或者 if i++ > 0 {}
	// go没有--i和++i

	// 不使用临时变量达到交换数据
	a := 2
	b := 3
	// 方法一：
	a = a + b
	b = a - b
	a = a- b
	fmt.Println(a, b)
	// 方法二
	a, b = b, a
	fmt.Println(a, b)
}