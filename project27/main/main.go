package main

import (
	"fmt"
	"math/rand"
	"time"
)

func test01() (newArr [26]byte){
	for i:=0; i < 26; i++ {
		newArr[i] = 'A' + byte(i) // 利用字符相加的特性，即ASCII码相加
	}
	return
}

// 求数组最大值和其下标
func test02(arr [5]int) (max int, index int) {
	max = 0
	index = 0
	for i, v := range arr {
		if v > max {
			max = v
			index = i
		}
	}
	return
}

// 随机生成一个数组，并将其反转
func test03() (arr [5]int) {
	rand.Seed(time.Now().UnixNano())		// 纳秒级
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)		// 0 <= n < 100
	}
	//输出反转前的数组
	fmt.Println("反转前：", arr)
	// 反转
	for i := 0; i < len(arr) / 2; i++ {
		arr[i], arr[len(arr) - 1 - i] = arr[len(arr) - 1 - i], arr[i]
	}
	// 输出反转后的数组
	fmt.Println("反转后：", arr)
	return
}


func main()  {
	// var arr [26]byte
	// arr = test01()
	// fmt.Println(arr)
	// for _, value := range arr {
	// 	fmt.Printf("%c, ", value)
	// }

	// arr02 := [...]int{-1,-8,2,8,10}
	// max, index := test02(arr02)
	// fmt.Println(max, index)

	// 随机种子
	// rand.Seed(time.Now().Unix())	// 精确到秒
	// rand.Seed(time.Now().UnixNano())	// 精确到纳秒

	_ = test03()
	
}