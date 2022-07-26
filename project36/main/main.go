package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	// map排序

	// 随机种子
	rand.Seed(time.Now().UnixNano())
	dict := make(map[int]int)

	for i := 1; i <= 10; i++ {
		dict[i] = rand.Intn(100)
	}
	fmt.Println(dict) // 整个显示出来是有序的

	// 遍历 只能用for range
	for k, v := range dict {
		fmt.Println(k, v) // 输出是无序的
	}

	// 如果按照map的key的顺序进行排序输出
	// 1. 先将map的key放到切片中
	// 2. 对切片进行排序
	// 3. 遍历切片，然后按照key来输出map的值

	var keys []int
	for k, _ := range dict { // for会将map的顺序打乱
		keys = append(keys, k)
	}
	sort.Ints(keys) // sort是在原切片上进行排序，覆盖原来那个切片
	fmt.Println(keys)

	for _, v := range keys {
		fmt.Println(v, dict[v])
	}

}
