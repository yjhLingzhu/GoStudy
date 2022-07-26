package main

import (
	"fmt"
)

func main() {
	fmt.Println(123)
	// 正数的原码，反码，补码都一样
	// 1 ==> 原码 [0000 0001] 反码 [0000 0001]  补码 [0000 0001]   最开始的那个数是符号为，正数的符号位是0，负数的符号位是1
	// -1 ==> 原码 [1000 0001] 反码 [1111 1110] 补码 [1111 1111]
	// 负数的反码=它的原码符号位不变，其他位取反
	// 负数的补码=它的反码+1（在最后一位开始加，逢2进1）
	// 0的反码补码都是0
	// 在计算机运行的时候都是以补码的形式进行的
	fmt.Println(2 & 3)
	fmt.Println(2 | 3)
	fmt.Println(2 ^ 3)

	// switch 语句用法
	var key byte
	fmt.Println("请输入一个字符：")
	fmt.Scanf("%c", &key)

	switch key {
	case 'a':
		fmt.Println("今天星期一")
	case 'b':
		fmt.Println("今天星期二")
	case 'c':
		fmt.Println("今天星期三")
	case 'd':
		fmt.Println("今天星期四")
	case 'e':
		fmt.Println("今天星期五")
	case 'f':
		fmt.Println("今天星期六")
	}
}
