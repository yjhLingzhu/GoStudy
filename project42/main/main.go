package main

import "fmt"

type integer int // 取别名

func (i integer) add() {
	fmt.Println("你好呀")
}

// 改变i的值
func (i *integer) change() { // i *integer 和 i integer 都不影响i.change()的调用，底层做了优化
	*i = *i + 1
}

type Person struct {
	Name string
}

func (person *Person) String() string {
	fmt.Println("123", person.Name)
	return person.Name
}

func main() {
	// 只要是数据类型(不过内置的数据类型需要取别名才能有方法)，它都可以有方法
	var i integer
	i.add()
	var j int
	c := integer(j) // 将j强制转换成integer类型后，它就具备integer的所有东西
	c.add()

	var i1 integer = 10
	i1.change()
	fmt.Println(i1)

	// 如果一个类型实现了String()这个方法，那么fmt.Println默认会调用这个变量的
	// String()进行输出。这个相当于python中的__str__魔法函数一样。
	var p1 Person = Person{"yjh"}
	// p1.String()  // 自己调用不需要传地址，底层优化了
	fmt.Println(&p1) // 结合Println()需要传地址
}
