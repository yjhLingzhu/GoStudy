package main


import (
	"fmt"
)


func main() {
	// 切片 var a []int   切片是引用类型，可以理解为python里面的列表

	// 切片定义方式一
	// 定义一个数组
	// arr01 := [...]int{1,2,3,4,5}
	// // 定义一个切片
	// slice := arr01[:3]  // 直接赋值的不需要make
	// slice[0] = 100		// 它是引用了arr01这个数组的前三个，所以在对slice的一个值复制的话，同时也会修改arr01的第一个值
	// fmt.Println(slice, len(slice), arr01)
	// fmt.Println("容量：", cap(slice))  // cap是获取容量的内置函数，它这个容量会随着你切片的个数增加不断变大
	// fmt.Printf("slice的首地址：%p, slice的地址：%p", &slice[0], &slice)

	// 切片定义方式二(推荐)
	// 通过 make创建一个切片，对于切片，必须make才能使用，不能var slice02 []int 这样直接使用
	slice02 := make([]int, 4, 10)	// 第一个是切片的类型，第二个是切片的长度，第三个是切片的容量，一般容量是长度的2倍，切片元素的默认值为0, 容量是可选参数
	fmt.Println(slice02)

	// 切片定义方式三
	var slice03 []string = []string{"yjh", "lz", "fzj"}
	fmt.Println(slice03, len(slice03), cap(slice03))

}