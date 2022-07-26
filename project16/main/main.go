package main

import (
	"fmt"
)

func add(n1 int, n2 int) {
	fmt.Println(n1, "+", n2, "=", n1 + n2)
}

// 函数作为参数传进去，类型不能带形参哟
func myFunc(funcvar func(int, int), n1 int, n2 int)  {
	funcvar(n1, n2)
}

// 例子： type myFunc func(int, int) int // myFunc 就等价func(int, int) int函数类型来使用了
type myFuncType func(int, int) // 相当于取别名，但是它俩又确实不是同一个数据类型，用相加的方法就可以证明，因为会报错
func myFunc2(funcvar myFuncType, n1 int, n2 int)  {
	funcvar(n1, n2)
}

// 函数可以对返回值直接命名
func cal(n1 int, n2 int) (sum int, sub int) {	// 这里的sum, sub就是对返回值直接命名了
	sum = n1 + n2
	sub = n1 - n2
	return			// 这样的话这里就不用写return sum, sub 这个语句了，因为前面定义了返回值，它会按照定义的顺序返回
}

func main()  {
	myFunc(add, 1,2)

	// 自定义数据类型：  type 自定义数据类型名 数据类型名  。 相当于给基本数据类型取了个别名
	// 例子： type myInt int // 这时myInt 就等价int 来使用了, 但是实际上它俩是不同类型，因为不能相加，会报错
	// 例子： type myFunc func(int, int) int // myFunc 就等价func(int, int) int函数类型来使用了
	type myInt int
	var a myInt
	a = 40
	fmt.Println(a)

	myFunc2(add, 2,2)

	// 函数可以对返回值直接命名
	res1, res2 := cal(4, 1)
	fmt.Println("res1=", res1, "res2=", res2)
}