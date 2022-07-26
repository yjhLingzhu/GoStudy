package main

import (
	"fmt"
	"reflect"
)

// 专门处理反射
func reflectTest01(b interface{}) {
	// 通过反射获取传入变量的type, kind，值
	// 1. 获取reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	// 2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)
	n2 := rVal.Int() + 2
	fmt.Println("n2=", n2)

	// 下面将rVal转成interface{}
	iv := rVal.Interface()
	// 将interface{} 通过断言转成需要的类型
	num2 := iv.(int)
	fmt.Println(num2)

}

func main() {
	// 演示对基本数据类型、interface{} 、reflect.Value进行反射的基本操作

	var num int = 100
	reflectTest01(num)
}
