package main

import (
	"fmt"
	"reflect"
)

// 使用反射来遍历结构体的字段，调整结构体的方法，并获取结构体标签的值

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

// 显示s的值
func (s Monster) Print() {
	fmt.Println("---start----")
	fmt.Println(s)
	fmt.Println("---end---")
}

// 返回两个数的和
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

// 接受四个值，个s赋值
func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

// 处理反射
func TestReflect(a interface{}) {
	// 获取reflect.Type
	rType := reflect.TypeOf(a)
	// 获取reflect.Value
	rVal := reflect.ValueOf(a)
	// 获取类别
	kd := rVal.Kind()
	// 如果不是结构体就退出
	if kd != reflect.Struct {
		fmt.Println("不是结构体....")
		return
	}

	// 获取结构体的字段数
	num := rVal.NumField()

	// 遍历结构体字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为=%v\n", i, rVal.Field(i))
		// 获取标签，需要通过reflect.Type来获取tag的值
		tagVal := rType.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d: tage=%v\n", i, tagVal)
		}
	}

	// 获取方法数量
	numOfMethod := rVal.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
	// 调用第二个方法，Call是调用该方法，
	// 我们的第二个方法是GetSum, 但是实际上在进行函数排序的时候它是按照ASCII码进行排序的，
	// 因此是GetSum > Print > Set 这样的顺序，所以这里调用的是Print这个方法，因此不需要参数
	rVal.Method(1).Call(nil)

	// 调用结构体的第一个方法Method(0)
	var params []reflect.Value // 声明了一个 []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	res := rVal.Method(0).Call(params) // 传入的参数是 []reflect.Value
	fmt.Println("res=", res[0].Int())  // 返回的结果是 []reflect.Value
}

func main() {
	var monster = Monster{
		Name:  "dazhu",
		Age:   18,
		Sex:   "female",
		Score: 95.5,
	}

	TestReflect(monster)
}
