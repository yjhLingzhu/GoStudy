package main

import (
	"fmt"
)

// 二分查找

func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findValue int) {

	if leftIndex > rightIndex {
		fmt.Println("没有找到!")
		return
	}

	middle := (leftIndex + rightIndex) / 2
	if (*arr)[middle] > findValue {
		BinaryFind(arr, leftIndex, middle - 1, findValue)
	} else if (*arr)[middle] < findValue {
		BinaryFind(arr, middle + 1, rightIndex, findValue)
	} else {
		fmt.Println("找到了，下标为: ", middle)
	}
}

func main()  {
	arr := [...]int{1,15,34,567,2567,9999}
	BinaryFind(&arr, 0, len(arr)-1, -567)

	
}