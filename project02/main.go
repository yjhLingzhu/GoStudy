package main

import "fmt"

func main() {
	n1, name, i := 10.1, "yjh", 100
	fmt.Println(n1, name, i)

	var c1 int = '北'		// 左右两边的类型没有保持一致，下面输出的时候需要用Printf输出，不能使用Println输出
	fmt.Printf("c1=%c c1的码值为：%d\n", c1, c1)	// 想要使用%d、%c等这些指定输出类型的话，需要使用Printf这个输出函数

	var c2 byte = 'a'		// 左右两边的类型保持一致，可以使用Println输出，不能使用Printf输出
	var c3 byte = '1'
	fmt.Println("c2=", c2)
	fmt.Println("c3=", c3)
}