package main

import (
	"fmt"
)

func test(value int) bool {
	if value <= 3 {
		return true
	} else {
		return false
	}
}

func main()  {
	// for循环 一
	for i := 1; i <= 5; i++ {
		fmt.Println("你好..")
	}

	// 二
	var j = 1	// 在外面定义了，for里面就可以不写
	for ; test(j); j++ {	// 循环条件可以是函数，只要返回值是bool类型的就行
		fmt.Println("你好j")
	}

	// 三
	var k = 1	// 循环变量的初始化
	for k <= 3 {	// 循环条件
		fmt.Println("你好k")	// 循环体
		k++		// 循环变量迭代
	}

	// 四 配合break使用，没有break的话就是死循环	相当于其他语言的while
	z := 1
	for {	// 等价与  for ; ; {}
		if z <= 3 {
			fmt.Println("break...")
			z++
		} else {
			break
		}
		
	}

	// for 循环遍历字符串-传统方式（字符串里面不能含有中文，否则会乱码，因为go是按utf编码的，一个中文占3字节）
	// var str string = "hello, world"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%c",str[i])
	// }

	var str string = "hello, world！北京"
	str2 := []rune(str)		// 切片   将str转成 []rune 这中类型
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c",str2[i])
	}
	fmt.Println()
	
	// for 循环遍历字符串-for-range的方式		和python的enumerate这个函数一样的 // 这种方式不会出现乱码的情况
	// str = "abc~ok"
	// for _, value := range str {
	// 	fmt.Printf("%c", value)
	// }
	// fmt.Println()

	// for 循环遍历字符串-for-range的方式		和python的enumerate这个函数一样的 // 这种方式不会出现乱码的情况
	// for-range 按照字符的方式遍历，有中文也是ok的
	str = "abc~ok！上海"
	for _, value := range str {
		fmt.Printf("%c", value)
	}

}