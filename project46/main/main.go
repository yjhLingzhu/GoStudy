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

// Working Computer
// 编写一个Working方法，接收一个Usb接口类型变量
// 只要是实现了 Usb接口（所谓实现Usb接口，就是指实现了Usb接口声明的所有方法）
func (computer *Computer) Working(usb Usb) {
	// 通过接口来调用Start和Stop
	usb.Start()
	usb.Stop()
}

func main() {
	// 接口
	// 测试
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}
	computer.Working(&phone)
	computer.Working(&camera)
}
