package main

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type pp Point // 取别名 但是pp 一个新的数据类型且可以和Point相互转换，但不可以直接相互赋值

type Monster struct {
	Name  string `json:"name"` // 这样加一个tag的话，在进行json序列话时会把字段改成对应的tag字段，注意冒号那里不能有空格
	Age   int    `json:"age"`  // `json:"age"`这就是结构体的tag
	Skill string `json:"skill"`
}

func main() {
	// 结构体字段在内存中是连续分布的
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	fmt.Println(&r1.leftUp.x, &r1.leftUp.y,
		&r1.rightDown.x, &r1.rightDown.y)

	// tag反射机制
	monster := Monster{"yjh", 5000, "芭蕉扇"}
	// 序列号json
	jsonStr, err := json.Marshal(monster) // 返回的是[]byte切片
	if err != nil {
		fmt.Println("序列号错误")
	}
	fmt.Println(monster)
	fmt.Println(jsonStr) // []byte切片
	// json.Marshal(monster)这个函数使用到了反射
	fmt.Println(string(jsonStr)) // 对[]byte切片强转成字符串
}
