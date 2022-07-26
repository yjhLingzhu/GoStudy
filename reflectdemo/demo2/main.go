package main

import (
	"fmt"
	"reflect"
)

// 专门处理反射
func reflectTest02(b interface{}) {
	// 通过反射获取传入变量的type, kind，值
	// 1. 获取reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	// 2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Println("rVal=", rVal)

	// 下面将rVal转成interface{}
	iv := rVal.Interface()
	fmt.Println("iv=", iv)
	// 将interface{} 通过断言转成需要的类型
	stu := iv.(Student)
	fmt.Println(stu.Age, stu.Name)

}

type Student struct {
	Age  int
	Name string
}

func main() {
	// 演示对基本数据类型、interface{} 、reflect.Value进行反射的基本操作

	var stu = Student{
		Name: "yjh",
		Age:  18,
	}
	reflectTest02(stu)
}
