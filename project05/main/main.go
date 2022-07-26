package main

import (
	"fmt"
	// 导包，这里要注意了，由于我的环境变量GOPATH 的路径是D:\workspace\owner\go\studyPrograms，所以如果想在我本机运行别的项目或者新建新的项目，
	// 那么需要将这个项目放到GOPATH这个路径下的src下面，因为导包是根据GOPATH去找的，然后自动拼接src，所以我们只需要写的是src后面那些路径即可。
	"go_code/project05/model"
)

// 1. 写一个程序， 获取一个变量的num的地址，并显示到终端
// 2. 将num的地址赋给指针ptr，并通过ptr去修改num的值

func main() {
	var num1 int = 300
	var ptr *int = &num1
	fmt.Printf("num1的地址：%v\n", ptr)

	*ptr = 250
	fmt.Printf("num1的值为：%v\n", num1)

	fmt.Println(model.HeroName)
	model.Name1()

}
