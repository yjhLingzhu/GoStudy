package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'y'
	var str string // 空字符串

	// 使用fmt.Sprintf转字符串
	str = fmt.Sprintf("%d", num1)
	fmt.Println(num1, num2, b, myChar, str)
	fmt.Printf("str type is %T, value is %q\n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type is %T, value is %q\n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type is %T, value is %q\n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type is %T, value is %q\n", str, str)

	// 使用 strconv 转字符串
	var num3 int = 99
	var num4 float64 = 23.456
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type is %T, value is %q\n", str, str)

	// f: 输出的格式  10: 小数保留10位, 64：num4的数据类型是64位
	str = strconv.FormatFloat(float64(num4), 'f', 10, 64)
	fmt.Printf("str type is %T, value is %q\n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type is %T, value is %q\n", str, str)

}
