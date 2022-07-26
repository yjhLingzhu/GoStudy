package main

import "fmt"

// AInterface 接口
type AInterface interface {
	Say()
}

// A 结构体
type A struct {
	Name string
}

// Say 实现了接口AInterface
func (a *A) Say() {
	fmt.Println("实现了接口...")
}

func main() {
	// 接口不能实例化，但是它可以指向实现了该接口的结构体的实例
	var ai AInterface // 接口类型的变量
	var a A           // 实现了接口类型的结构体变量
	ai = &a           // 要给地址
	ai.Say()
}
