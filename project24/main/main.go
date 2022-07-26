package main

import (
	"fmt"
)

func sum()  {
	var intArr [5]int32
	intArr[0] = 1
	intArr[1] = 1
	intArr[2] = 1
	intArr[3] = 1
	intArr[4] = 1
	var sum_ int32
	sum_ = 0
	// fmt.Println(len(intArr))
	for i := 0; i < len(intArr); i++ {
		sum_ += intArr[i]
	}
	fmt.Println("sum=", sum_)
	fmt.Println(intArr)
}


func main()  {
	sum()
}