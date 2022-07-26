package main


import (
	"fmt"
)

func main()  {
	// 切片就是能动态增加个数
	// 声明一个map切片
	data := make([]map[string]string, 0)
	// 将数据放进去
	temp := map[string]string{"name": "yjh", "age": "18",}
	data = append(data, temp) // append一定要有一个变量来接收的
	// data[0] = temp
	temp = map[string]string{"name": "fzj", "age": "19", "addr": "广东省"}
	data = append(data, temp)
	fmt.Println(data)
	// 这种在web作为返回的json串很实用
}