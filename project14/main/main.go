package main


import (
	"fmt"
	"go_code/project14/db"
	"go_code/project14/utils"		// 同一个文件夹下的所有文件不能存在多个不同的包。且这文件夹下的所有文件不能存在相同的函数或全局变量名
)

func main()  {
	a := utils.Add()
	b := db.Add1()
	fmt.Println(a)
	fmt.Println(b)
}