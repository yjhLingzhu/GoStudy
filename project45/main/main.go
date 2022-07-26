package main

import "fmt"

// Goods 基本语法
type Goods struct {
	Name  string
	Price int
}

// Book 结构体
type Book struct {
	Goods  // 这里就是嵌套匿名结构体Goods
	Writer string
	int    // 匿名字段
}

func main() {
	// 继承
	/*
		1. 在结构体中嵌入一个匿名结构体就是继承了
		2. 结构体可以使用嵌套匿名结构体所有的字段和方法，不管首字母大写还是小写
		3. 当结构体和匿名结构体有相同字段或方法时，采用就近原则访问，如果希望
		   访问匿名结构体里面的字段或者方法的话，可以通过匿名结构体名来区分，即
		   book.Goods.Name这样访问，不要book.Name这样访问
		   （这里的就近原则指的是当前查找到哪个层级了，即哪个结构体了）
		4. 要想知道当前在哪个层级，主要看当前代码的receiver是哪个结构体
		5. 如果在嵌入多个匿名结构体时，且这些匿名结构体都具有某些相同的字段
		   或者方法（这些字段和方法在这个结构体没有）时，
		   那么在访问这些字段或者方法时就要指明结构体名称，否则会报错
		6. 如果struct嵌套的是一个有名字的结构体，这种模式就是组合，这样的话
		   在访问时就必须指定结构体名字
	*/
	// 使用
	book := &Book{}
	book.Goods.Name = "yjh" // 等价 book.Name  相当于在本结构体找不到时，它会找继承的结构体中有没有这个属性或者方法
	book.Goods.Price = 188  // 等价 book.Price
	book.Writer = "fzj"
	book.int = 20 // 匿名字段的用法
	fmt.Println(*book)
}
