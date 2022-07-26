package main

import (
	"fmt"
)

func main()  {
	fmt.Println(123)
	// 获取控制台输入的数据
	
	// var age byte = 10
	// if age > 1 {
	// 	fmt.Println(5)
	// }

	// fmt.Scanln  获取一行数据，以回车为获取标识
	// fmt.Scanf 默认以空格为标识，可以指定获取标识
	var name string
	var age byte
	var sal float32
	var isPass bool
	
	// 方式一
	// fmt.Println("请输入姓名：")
	// fmt.Scanln(&name)

	// fmt.Println("请输入年龄：")
	// fmt.Scanln(&age)

	// fmt.Println("请输入薪水：")
	// fmt.Scanln(&sal)

	// fmt.Println("是否通过：")
	// fmt.Scanln(&isPass)

	// fmt.Println("name=", name, "age=", age, "sal=", sal, "isPass=", isPass)

	// 方式二
	fmt.Println("请输入姓名，年龄，薪水，是否通过，用空格隔开：")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Println("name=", name, "age=", age, "sal=", sal, "isPass=", isPass)
}