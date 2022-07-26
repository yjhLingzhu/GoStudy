package main

import "fmt"

// 结构体是值拷贝
type A struct {
	num int
}

// 方法  这里的方法和函数是有点区别的 方法多了给指定是哪个数据类型才能调用
func (a A) test() { // 这里的(a A)是指定这个test()函数只能被A类型的变量调用，即给A绑定了一个方法test()
	fmt.Println(a.num)
}

// 其实可以理解为go在没有class的情况下，它是将class里面的函数独立出去了而已
// 结构体 + 这个结构体的所有函数 = class

// 完整的方法声明
// func (receciver type) methodName(args... int) (返回的参数列表){
//		方法体
// 		return 返回值
//}

func main() {
	var t A // 声明一个变量
	t.num = 100
	t.test() // 这个test() 是一个方法，而且是只能给A这种数据类型调用的方法
	// 它这样调用了就相当于把t传给了func (a A) test() {}中的形参a
	// 所以在函数体内使用的形参a就是实例t
}
