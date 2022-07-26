package main

import (
	"fmt"
)

func main() {
	var dict = make(map[string]string)
	dict["name"] = "yjh"
	dict["age"] = "18"
	dict["addr"] = "广东省"
	fmt.Println(dict)

	// 查找
	val, ok := dict["name"]
	if ok {
		fmt.Println("有这个key, 值为：", val)
	} else {
		fmt.Println("没有这个key")
	}

	// 删除
	// delete(dict, "addr") // 就算key不存在，它也不会报错
	// fmt.Println(dict)

	// 遍历 只能用for range
	for k, v := range dict {
		fmt.Println(k, v)
	}

	// 长度
	fmt.Println(len(dict))

}
