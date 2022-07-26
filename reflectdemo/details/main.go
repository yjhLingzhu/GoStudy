package main

import (
	"fmt"
	"reflect"
)

// 通过反射修改 num int的值，修改student的值
func reflect01(b interface{}) {
	// 获取反射的值
	rVal := reflect.ValueOf(b)
	// 这里传过来的b是指针类型的话，需要使用Elem()这个方法
	// 转成非指针类型，然后再可以使用Value绑定的方法
	// 这个Elem类似于*p这样的操作。
	// Set开头的函数前面一定要Elem来调用，而Elem前面一定要指针
	// 类型来调用。
	rVal.Elem().SetInt(20) // 给该变量赋值为20
}

func main() {
	var num int = 10
	reflect01(&num) // 传地址过去
	fmt.Println(num)
}
