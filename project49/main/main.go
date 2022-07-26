package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

// 使用系统排序方法对结构体进行排序
// 1. 声明一个结构体
type Hero struct {
	Name string
	Age  int
}

// 2. 声明一个Hero结构体切片类型
type HeroSlice []Hero

// 3. 实现Interface 接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less 方法就是决定你使用什么标准进行排序
// 1. 按照Hero的年龄从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	// 这里对年龄进行排序
	// return hs[i].Age < hs[j].Age
	// 按照名字进行排序
	return hs[i].Name < hs[j].Name
}

func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}
func main() {
	// 测试
	rand.Seed(time.Now().UnixNano()) // 不能放到for里面，否则值不变
	var heroes HeroSlice             // 定义一个列表，里面每个元素是Hero类型
	// heroes = make(HeroSlice, 0)
	hero := Hero{} // 定义一个dict
	for i := 0; i < 10; i++ {
		name := "yjh_" + strconv.Itoa(rand.Intn(100))
		age := rand.Intn(7) + 18
		hero.Age = age
		hero.Name = name
		heroes = append(heroes, hero)
	}
	fmt.Println("排序前：", heroes)
	sort.Sort(heroes)
	fmt.Println("排序后：", heroes)
}
