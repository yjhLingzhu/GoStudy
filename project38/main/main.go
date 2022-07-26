package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// 定义猫的结构体
type Cat struct {
	Name  string
	Age   int
	Color string
}

// 传统方法实现
func traditon() {
	cats := make(map[string]map[string]string)
	// 定义随机种子
	rand.Seed(time.Now().UnixNano())

	catName := []string{"小米", "大米"}

	for _, value := range catName {
		age := strconv.Itoa(rand.Intn(5) + 1)
		weight := strconv.Itoa(rand.Intn(15) + 10)
		cats[value] = map[string]string{"age": age, "weight": weight}
	}

	var cat string
	fmt.Println("请输入小猫的名称:")
	fmt.Scanf("%s", &cat)

	val, ok := cats[cat]
	if ok {
		fmt.Printf("这个猫的年龄是:%v，体重是%v", val["age"], val["weight"])

	} else {
		fmt.Println("不存在这个猫")
	}
}

func main() {
	// 随机生成两个猫的信息，控制台输入存在的名字则将这些信息输出否则提示这个猫不存在

	// 传统方法
	// traditon()

	// 使用结构体的方法实现
	var cat1 Cat
	cat1.Name = "fzj"
	cat1.Age = 18
	cat1.Color = "白色"
	fmt.Println(cat1)

}
