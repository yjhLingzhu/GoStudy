package main

import (
	"fmt"
	"errors"
)

// 使用defer + recover 来捕获和处理异常
func test()  {
	defer func ()  {
		err := recover() // recover() 是内置函数，可以捕获到异常
		if err != nil {
			fmt.Println("err=", err)
		}
	}()


	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("结果：", res)
}

func test1() (n int) {
	n = 1 + 1
	return		// 这里不写返回值，默认返回(n int)中的n，如果写了就返回写的那个值
}

// 函数去读配置文件
// 如果文件名传入不正确，返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		return errors.New("读取文件错误...")
	}
}

func test2()  {
	err := readConf("config.ini")
	if err != nil {
		// 如果失败，输出错误并终止程序
		panic(err)
	}
	fmt.Println("继续执行....")
}

func main()  {
	test()
	// n := test1()
	// fmt.Println("n=", n)
	test2()
	fmt.Println("main()的内容... ")
}