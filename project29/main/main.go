package main

import (
	"fmt"
)

func main()  {
	// 用append内置函数可以对切片进行动态追加

	arr01 := []int{1,2,3}
	fmt.Println("arr01: ", arr01)
	// 追加一个数
	arr01 = append(arr01, 4)  // 这里要用一个变量来接受，否则报错，其实其内部是新建了一个新的切片的，arr01本身是没有变化的，我们这里赋值给了arr01才导致其发生变化，其实arr01可以换成别的切片变量
	// 一次性追加多个
	arr01 = append(arr01, 5,6,7)
	fmt.Println("追加后：", arr01)

	// 通过append将切片追加进去
	arr01 = append(arr01, arr01...)
	fmt.Println("追加后：", arr01)

	// 拷贝操作
	a := []int{1,2,3,4,5}
	slice01 := make([]int, 10)	// 长度一定要大于等于a的长度，不然拷贝后是空值
	// 将a拷贝到slice01中去
	copy(slice01, a) // 是切片类型才能执行拷贝操作
	// slice01 和 a 的数据空间是独立的  
	fmt.Println("拷贝后：", slice01)
	// 输出是： [1 2 3 4 5 0 0 0 0 0]
}