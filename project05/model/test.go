package model

import (
	"fmt"
)

var heroName string  = "你呀的"

// go通过标识符的第一个字母大小写的区分来达到包的作用域效果，小写开头的，则只能在本包使用，不需要导包就可以使用；
// 大写开头的可以在别的包引用，需要导包, 在本包使用也不需要导包
// 可以这样理解：
//     一个包下面的所有函数和变量都不能重名，理解为它内部编译时是将这个包下所有go文件的内容整合到了一个很大的go文件中。
func Name1()  {
	fmt.Println(563)
	name()
	fmt.Println(HeroName)
}
