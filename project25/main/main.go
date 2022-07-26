package main

import (
	"fmt"
)

func main()  {
	// 初始化数组的四种方式
	// 一
	var array1 [3]int = [3]int{1,2,3}
	fmt.Println(array1)

	// 二
	var array2 = [3]int{4,5,6}
	fmt.Println(array2)

	// 三
	var array3 = [...]int{7,8,9}
	fmt.Println(array3)

	// 四， 这里指定每个下标和其对应的值
	var array4 = [...]int{1:800, 0:900, 2:700}
	fmt.Println(array4)
}