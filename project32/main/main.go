package main

import (
	"fmt"
)

func main()  {
	// 二维数组

	var arr [4][6]int

	// len(arr)获取的是行数， len(arr[0])获取的是列数
	fmt.Println(len(arr), len(arr[0]))

	// 二维数组的初始化
	var arr02 = [2][]int{{1,2,3}, {4,5}} // 和python一样，可以定义这种不等长度的二位数组

	for i := 0; i < len(arr02); i++ {
		for j := 0; j < len(arr02[i]); j++ {
			fmt.Println(arr02[i][j])
		}
	}

	for i, v1 := range arr02 {
		for j, v2 := range v1 {
			fmt.Printf("arr02[%v][%v]=%v ", i, j, v2)
		}
		fmt.Println()
	}
	
}