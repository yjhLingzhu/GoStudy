package main

import "fmt"

// Usb 定义一个接口 不能包含任何变量
type Usb interface {
	// 声明了两个没有实现的方法
	Start() // method(参数列表) 返回值列表
	Stop()
}

// Phone 结构体
type Phone struct {
}

// Start 方法
func (phone *Phone) Start() {
	fmt.Println("phone start....")
}

// Stop 方法
func (phone *Phone) Stop() {
	fmt.Println("phone stop....")
}

// Camera 结构体
type Camera struct {
}

// Start Camera方法
func (camera *Camera) Start() {
	fmt.Println("camera start....")
}

// Stop Camera方法
func (camera *Camera) Stop() {
	fmt.Println("camera stop....")
}

// Computer 结构体
type Computer struct {
}

func main() {
	// 接口是对继承的扩展
	// 继承的价值主要在于：解决代码的复用性和可维护性
	// 接口的价值主要在于：设计好各种规范(方法)，让其他自定义类型去实现这些方法
	// 接口比继承更加灵活
	// 继承是满足is - a 的关系，而接口只需满足like - a的关系
	// 接口在一定成都上实现代码解耦

	// 多态：不同的实例调用相同的方法，呈现不同的内容（前面usb那个例子就是多态）
	// 定义一个Usb接口数组，可以存放Phone和Camera的结构体
	// 这里就体现了多态，同时这个数组可以正真称得上是python里面的列表了，
	// 因为它里面的元素是不同类型的（从结构体角度看，如果从接口角度看，它们都属于usb类型）
	var usbArr [3]Usb
	fmt.Println(usbArr)
	usbArr[0] = &Phone{} // 这里传地址是因为是实现接口时绑定结构体用的是指针类型
	usbArr[1] = &Phone{}
	usbArr[2] = &Camera{}
	fmt.Println(usbArr)
}
