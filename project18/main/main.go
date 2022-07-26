package main

// 先执行全局变量定义，再执行init函数，再执行main函数

// 同一个包下的go文件是可以每个文件都有自己的init函数的，它不会被纳进重名检测中
import (
	"fmt"
	"go_code/project18/utils"
)

func main()  {
	fmt.Println("Age=", utils.Age, "Name=", utils.Name)
	fmt.Println("Addr=", utils.Addr)
}