package main


import (
	"fmt"
)

func sum(n1, n2 int) int {
	// 程序执行到defer时，会将defer的代码压入到一个defer栈中，等到函数执行完毕后，它再从栈中依次执行这些defer的代码，顺序是栈的顺序（先入后出）
	defer fmt.Println("ok1")  // 3
	defer fmt.Println("ok2")  // 2

	n := n1 + n2
	fmt.Println("n=", n)	// 1
	return n
}

// Name := "yjh"		// 全局变量不能执行赋值语句，因为这个语句等价于 var Name string    Name = "yjh"
// 函数外只能执行 变量定义，如 var Name string 或变量初始化，如 var Name = "yjh"

// go里 一个中文字符是占用3个字节，其他字符占1个字节
// go中的len()函数是返回字节的长度，不像python是返回字符的长度

func main()  {
	res := sum(1,2)
	fmt.Println("res=", res)  // 4
}