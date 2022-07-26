package main


import (
	"fmt"
)

var (
	// Fun1是全局匿名函数，因为将一个匿名函数赋值给一个全局变量
	Fun1 = func (n1, n2 int) int {
		return n1 * n2
	}
)

func main()  {
	// 匿名函数

	n1 := 2
	n2 := 3
	res1 := func (n1, n2 int) int {
		return n1 + n2
	}(n1, n2)	// 定义完直接开始调用
	
	fmt.Println("res1=", res1)

	// 将匿名函数赋值给一个变量
	a := func (n1, n2 int) int {
		return n1 + n2
	}
	res2 := a(1,1)
	fmt.Println("res2=", res2)

	// 全局匿名函数
	res3 := Fun1(2,2)
	fmt.Println("res3=", res3)
}
	