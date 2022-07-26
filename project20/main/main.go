package main


import (
	"fmt"
	"strings"
)

// 累加器
func AddUpper() func (int) int {
	n := 10
	return func (x int) int {
		n = n + x
		return n
	}
}

// 后缀闭包
func makeSuffix(suffix string) func (string) string {

	return func (name string) string {
		if !strings.HasSuffix(name, suffix) {	// 不是指定后缀
			return name + suffix
		}
		return name
	}
	
}

func main()  {

	// 不是函数中套函数，不属于闭包，所以就算func1指向一个函数，在它调用的时候还是当一个新函数进栈处理
	func1 := func (n1 int) int {
		n2 := 10
		return n2 + n1
	}

	fmt.Println(func1(1)) // 11
	fmt.Println(func1(2)) // 12

	// 闭包
	func2 := AddUpper()
	fmt.Println(func2(1)) // 11
	fmt.Println(func2(2)) // 13	 这里输出13可以理解为func2指向了一个函数，再每次调用这个函数时，
						  // 它不会像调用新函数那样将函数进栈，而是直接去执行栈中的这个存在的函数(因为这个函数是在AddUpper函数里面，它区别与匿名函数赋值变量)，
						  // 所以这个函数的局部变量会一直累积上次调用的结果
	

	// 闭包就是匿名函数+它里面用了别人的局部变量(全局变量不算) 

	// 需求
	// 用闭包实现，传入一个名字，判断这个名字是否是指定后缀的，如果是则返回，否则加上指定后缀返回
	func3 := makeSuffix(".png")
	fmt.Println("文件结果：", func3("wind.png"))
	fmt.Println("文件结果：", func3("hahah"))

}