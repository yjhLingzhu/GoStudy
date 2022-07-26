package main

import (
	"fmt"
	"go_code/project44/model"
)

func main() {
	// 工厂模式

	stu := model.GetStudent("yjh", 18)
	fmt.Println(*stu)
}
