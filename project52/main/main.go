package main

import (
	"fmt"
	"math/rand"
	"time"
)

// TypeJudge items ...interface{}表示接受任意类型多个参数，相当于python的*args
func TypeJudge(items ...interface{}) {
	for _, v := range items {
		switch v.(type) {
		case bool:
			fmt.Println("bool")
		case float64:
			fmt.Println("float64")
		case int, int64:
			fmt.Println("int, int64")
		case nil:
			fmt.Println("nil")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("不确定")
		}
	}
}
func main() {
	// var n1 int = 10
	// var n2 float64 = 10.1
	// var n3 bool = false
	// var n4 string = "yjh"
	// TypeJudge(n1, n2, n3, n4)

	rand.Seed(time.Now().UnixNano())
	base_str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	var random string = ""
	for i := 0; i < 16; i++ {
		random += string(base_str[rand.Intn(len(base_str)-1)])
	}
	fmt.Println(len(base_str))
	fmt.Printf("%v", random)
}
