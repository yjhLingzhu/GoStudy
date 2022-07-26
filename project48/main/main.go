package main

import "fmt"

type BInterface interface {
	test01()
}

type CInterface interface {
	test02()
}

type AInterface interface {
	BInterface // 继承了接口
	CInterface // 这里继承的接口不能存在相同的方法名
	test03()
}

// 如果要实现AInterface接口，那么需要将BInterface、CInterface里面的所有方法
// 以及test03()都实现才行

type Stu struct {
}

func (stu *Stu) test01() {
	fmt.Println("test01...")
}
func (stu *Stu) test02() {
	fmt.Println("test02...")
}
func (stu *Stu) test03() {
	fmt.Println("test03...")
}

type NullInterface interface{}

type Te struct{} // 实现多个接口

func (te *Te) test01() {
	fmt.Println("Te test01...")
}

func (te *Te) test02() {
	fmt.Println("Te test02...")
}

func main() {
	var stu Stu
	var a AInterface // 接口是引用类型，
	a = &stu
	a.test01()
	a.test02()
	a.test03()

	var n NullInterface = stu // 我们可以把任何类型赋值给空接口
	fmt.Println(n)
	var t int = 10
	n = t
	fmt.Println(n)

	var te Te // 初始化方式一
	h := Te{} // 初始化方式二
	h.test01()
	te.test01()
	te.test02() // 实现了多个接口，所以它能调用那些接口的方法
	var bt1 BInterface = &te
	var ct1 CInterface = &te
	bt1.test01() // 虽然te是实现了多个接口，但是它给BInterface的bt1这个变量引用了，所有它只能调用这个接口里面的方法
	ct1.test02() //虽然te是实现了多个接口，但是它给CInterface的ct1这个变量引用了，所有它只能调用这个接口里面的方法
}
