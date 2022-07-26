package main

import "fmt"

// Point .
type Point struct {
	x int
	y int
}

func main() {
	// 空接口可以接受任意类型，但是不能把空接口变量赋值给任意类型变量
	var a interface{}
	var point = Point{1, 2}
	a = point // ok
	fmt.Println(a)

	var b Point
	// a = b // error
	b = a.(Point)  // ok,这是类型断言,这种要知道a之前是哪个类型的才行
	fmt.Println(b) // 如果想将一个类型赋值给空接口类型，那么要用类型断言，其实就是强制转换

	var c int
	var e float32
	e = 10.2
	c = int(e) // 普通强制转换
	fmt.Println(c)

	// 带检测的类型断言
	var x interface{}
	var b2 float32
	x = b2
	y, ok := x.(float64)
	if ok {
		fmt.Println(y)
	}

	// 等价
	// if y, ok := x.(float64); ok {
	// 	fmt.Println(y)
	// }
}
