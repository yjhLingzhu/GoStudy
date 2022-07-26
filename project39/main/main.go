package main

import "fmt"

// 结构体是值类型
type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	prt    *int
	slice  []int
	map1   map[string]string
}

type Dog struct {
	Name string
	Age  int
}

func main() {
	// 结构体里面字段的类型是：指针，slice，map时，它们的默认值是nil，即还没分配空间
	// 所以要先make才能使用
	var p1 Person // 结构体声明方式一
	arr := [...]int{1, 2, 1, 4, 5, 6}
	p1.slice = arr[:3]                // 相当于make了
	p1.slice[0] = 999                 // 对里面的元素进行修改时要确保已经make了
	p1.map1 = make(map[string]string) // 先make
	p1.map1["name"] = "yjh"
	fmt.Println(p1)

	var d1 Dog = Dog{}         // 结构体声明方式二 等价 p2 := Dog{}
	var d2 Dog = Dog{"fzj", 1} // 结构体声明方式二 等价 p2 := Dog{}
	fmt.Println(d1, d2)

	// 方式三
	d3 := new(Dog) // 等价 var d3 *Dog = new(Dog)   new返回的是指针
	(*d3).Age = 2
	(*d3).Name = "jzf"
	// 上面的赋值可以等价为：d3.Age = 2  d3.Name = "jzf"
	// 虽然d3是指针类型，但是你这样写了后，底层会对d3进行
	//取值运算, 即d3 -> (*d3)，所以这样写会成功(即使d3是指针)
	fmt.Println(*d3)
	d3.Age = 100
	d3.Name = "hhh"
	fmt.Println(*d3)

	// 方式四
	var d4 *Dog = &Dog{} // 因为d4是指针，所以要用取址符，new相当于在内部做了这样的事，将地址返回去了
	(*d4).Age = 3
	(*d4).Name = "yjh" // .的运算级比*高
	fmt.Println(*d4)
	d4.Age = 666
	d4.Name = "hhhwa"
	fmt.Println(*d4)
}
