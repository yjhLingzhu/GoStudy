package main


import (
	"fmt"
)

func main()  {
	// map 可以理解为python的dict
	// 声明的格式： var name map[keytype]valuetype
	// 声明一个map是不会分配内存的，初始化需要make，分配内存后才能赋值和使用
	// make的作用就是给map分配空间
	// 创建map方式一
	var a map[string]string
	a = make(map[string]string, 10) // 数据类型，空间大小
	a["name"] = "yjh"
	fmt.Println(a)

	// 创建map方式二（推荐）
	var dict = make(map[string]string)
	dict["name"] = "yjh"
	dict["age"] = "18"
	fmt.Println(dict)

	// 创建map方式三
	var dict1 = map[string]string{
		"name": "yjh",
	}
	dict1["age"] = "19"
	fmt.Println(dict1) 

	// var num = "10"
	// fmt.Printf("%T", num)

	// 案例
	stu := make(map[string]map[string]string)
	stu["stu01"] = make(map[string]string)
	stu["stu01"]["name"] = "yjh"
	stu["stu01"]["age"] = "18"
	stu["stu01"]["class"] = "一班"
	
	stu["stu02"] = make(map[string]string)
	stu["stu02"]["name"] = "fzj"
	stu["stu02"]["age"] = "19"
	stu["stu02"]["class"] = "二班"
	fmt.Println(stu)
}