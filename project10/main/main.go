package main

import (
	"fmt"
)

func main()  {
	// var n1 int32 = 10
	// var n2 int32

	switch n1 := 10; {
		case n1 > 50:
			fmt.Println(n1, "大于50")
			fallthrough		// 穿透，继续往下执行, 只穿透一层
		case n1 > 30:
			fmt.Println(n1, "大于30小于50")
		
	}

	var num1 int = 10
	switch num1 {
	case 10:
		fmt.Println("ok1")
	case 20:
		fmt.Println("ok2")
	default:
		fmt.Println("没有匹配...")
	}

	var score int = 100
	switch {
	case score > 90:
		fmt.Println("优秀")
	}
}