package main

import (
	"fmt"
)

func main()  {
	// string 底层是byte数组，因此string也可以进行切片处理
	str := "yjh@qq.com"
	slice := str[:3]
	fmt.Println(slice)

	// string是不可变的，str[0] = 'a' 这样不行，会报错

	// 如果想修改字符串，可以先将string -> []byte切片
	// 将yjh改成fzj
	arr01 := []byte(str)
	slice02 := []byte{'f','z','j'}
	copy(arr01, slice02)
	str = string(arr01) // 将[]byte切片转成字符串
	fmt.Println("修改后：", str)
	// 我们转成[]byte后可以处理英文和数字，但是不能处理中文
	// 因为中文是按三个字节处理的，因此会出现乱码
	// 解决办法： 将string转成[]rune即可，因为[]rune是按字符处理，兼容汉字
	 
	str1 := "yjh@qq.com"
	arr02 := []rune(str1)
	slice03 := []rune{'哈','哈','哈'}
	copy(arr02, slice03)
	str1 = string(arr02) // 将[]byte切片转成字符串
	fmt.Println("修改后-：", str1)
}