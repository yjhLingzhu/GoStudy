package main

import (
	"fmt"
)

func test01(arr [3]int)  {		// [3]int 是数据类型（长度是数据类的一部分），底层认为 [3]int 和 [4]int是不同类型的。所以形参中一定要写上长度
	arr[0] = 99
	fmt.Println(arr)
}

func test02(arr *[3]int)  {		// 传了地址过来，这里可以修改之前的数组
	(*arr)[0] = 100
	fmt.Println(*arr)
}


func main()  {
	// for range 遍历数组

	var heroes [3]string = [3]string{"yjh", "lz", "feizhuji"}
	for _, value := range heroes {
		fmt.Println(value)
	}

	// 数组属于值类型，在默认情况下是值传递，因此只会进行值拷贝。数组间不会相互影响。
	arr := [3]int{1,2,3}
	// test01(arr)
	// fmt.Println(arr)

	test02(&arr)
	fmt.Println(arr)
}