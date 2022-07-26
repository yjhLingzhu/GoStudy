package main

import (
	"fmt"
	"time"
)


func main()  {
	// 获取当前时间
	now := time.Now()
	fmt.Printf("now=%v, now type is %T\n", now, now)

	// 获取年月日时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	// 格式化日期  方式一
	fmt.Printf("当前年月日时分秒 %d-%d-%d %d:%d:%d\n", now.Year(),
	now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	dataStr := fmt.Sprintf("当前年月日时分秒 %d-%d-%d %d:%d:%d", now.Year(),
	now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	fmt.Printf("dataStr=%v\n", dataStr)

	// 方式二 2006-01-02 15:04:05 等价python的 %Y-%m-d% %H:%M:%S
	fmt.Printf(now.Format("2006-01-02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

	// 时间常量
	// 每个一秒输出一个数
	i := 0
	for {
		i++
		fmt.Println(i)
		// 休眠一秒
		// time.Sleep(time.Second)
		// 休眠0.1秒 time.Millisecond -> 1毫秒
		time.Sleep(time.Millisecond * 100)
		if i == 10 {
			break
		}
	}
}