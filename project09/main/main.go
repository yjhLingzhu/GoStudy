package main

import (
	"fmt"
)

func main()  {
	var age byte

	fmt.Println("请输入年龄：")
	fmt.Scanf("%d", &age)

	if age > 18 && age < 30{
		fmt.Println("年龄大于18，你输入的是:", age)
	} else {	// else 关键字不能换行，go强制要求，否则编译不通过
		fmt.Println("年龄小于18，你输入的是:", age)
	}
}