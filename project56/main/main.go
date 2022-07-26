package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// 用具有缓冲区的方式读取数据
	// newReaderFile()

	// 没有缓冲区
	// readerFile()

	// Open和OpenFile都是os里面打开文件的函数，但是OpenFile可用于不存在的路径
	// 因此可以这样说Open是打开一个文件，OpenFile是创建一个文件(存在就打开，不存在就创建)
	// demo1()

	// demo2()

	// demo3()

	// demo4()

	// demo5()

	// checkfolder()

	// copyFile()

	statis()
}

func newReaderFile() {
	// 打开文件
	f, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("打开文件失败！")
		return
	}
	defer f.Close()

	// 将文件和reader绑定，这个reader是有缓冲区的，可以适用大文件小文件
	reader := bufio.NewReader(f)
	// 这个缓冲区是使其在读取文件时，先从缓冲区读取到内存，如果缓存区没有数据了，那么它才会去磁盘中读取数据同时也更新数据到缓冲区
	for {
		str, err := reader.ReadString('\n') // 读取以‘\n’结尾的数据作为这次取得的数据
		if err == io.EOF {
			break
		}
		fmt.Print(str) // 这里不用换行，因为它也把'\n'读进去了
	}
	fmt.Println("文件读取完毕。")
}

// 使用ioutil.ReadFile一次性将文件读取到位（直接去磁盘读取，没有缓冲区）
func readerFile() {
	// 如果文件很大的话，会导致效率很低，所有适用于小文件
	content, err := ioutil.ReadFile("./test.txt") // 这个方法不需要显示打开和关闭文件，它自己在里面已经做了
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(content))
}

// 创建一个新的文件，写入5句“hello go”
func demo1() {
	filePath := "./demo1.txt"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 准备写入数据
	str := "Hello go word!\r\n"
	// 写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 因为writer是带缓存的，因此在调用WriteString写入数据时，它是将数据写到缓存上了，还没真正写到
	// 磁盘上面, 所以需要调用Flush方法，将缓存的数据真正写到磁盘文件中
	writer.Flush()

}

// 打开一个存在的文件，将原来的内容覆盖成新的内容10句“你好，yjh”
func demo2() {
	filePath := "./demo1.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666) // O_CREATE这个模式是覆盖的模式，这里可以只用O_CREATE也可以到达效果
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 创建一个带缓存的writer
	writer := bufio.NewWriter(file)
	// 写入的内容
	str := "你好，yjh\r\n"
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	// 调用Flush方法真正将数据写到磁盘上
	writer.Flush()
}

// 打开一个已经存在的文件，在后面追加内容 “哈哈哈”
func demo3() {
	// 打开文件
	filePath := "./demo1.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 使用带缓存的writer写入数据
	writer := bufio.NewWriter(file)
	// 写入的数据
	str := "哈哈哈哈哈\r\n"
	writer.WriteString(str) // 这里数据是写到了缓存
	// 将数据真正写到磁盘中
	writer.Flush()
}

// 打开一个存在的文件，将原来的内容读出显示在终端，并且追加5句“Hello，北京！”
func demo4() {
	// 打开文件
	filePath := "./demo1.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 将数据读取处理并显示到终端
	// 使用带缓存的reader
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF { //文件末尾
			break
		}
		fmt.Print(str) // 这里不用换行，因为它也把'\n'读进去了
	}
	// 使用带缓存的writer
	writer := bufio.NewWriter(file)
	// 写入的数据
	str := "Hello, 北京！\r\n"
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	// 将数据写到磁盘
	writer.Flush()
}

// 将一个文件的内容写入到另外一个文件，这两个文件已经存在了
// 使用ioutil.ReadFile/ioutil.WriteFile
func demo5() {
	// 打开文件
	src := "./test.txt"
	dst := "./demo1.txt"

	c1, err := ioutil.ReadFile(src)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(dst, c1, 0666) // 这样直接写是覆盖的方式写
	if err != nil {
		panic(err)
	}
}

// 判断文件或文件夹是否存在
func checkfolder() {
	filePath := "./test.txt"
	_, err := os.Stat(filePath)
	if err == nil { // 不报错，说明文件存在
		fmt.Println("文件存在！")
	} else if os.IsNotExist(err) { // 报这个错说明文件不存在
		fmt.Println("文件不存在1")
	} else {
		// 报其他错说明不确定文件是不是存在，这里统一认为不存在处理
		fmt.Println("文件不存在2")
	}
}

// 拷贝文件
func copyFile() {
	src := "../src/2.jpg"
	dst := "../dst/new.jpg"

	s_file, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer s_file.Close()
	// 新建一个reader
	reader := bufio.NewReader(s_file)

	d_file, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer d_file.Close()
	// 新建一个writer
	writer := bufio.NewWriter(d_file)

	n, err := io.Copy(writer, reader)
	if err != nil {
		panic(err)
	}
	fmt.Println("长度：", n)
}

// 定义一个结构体，用于保存统计结果
type CharCount struct {
	chCount    int // 记录英文个数
	NumCount   int // 记录数字的个数
	SpaceCount int // 记录空格的个数
	OtherCount int // 记录其他字符的个数
}

// 统计一个文件中含有的英文、数字、空格及其他字符数量
func statis() {
	// 思路：
	// 打开一个文件，创建一个reader
	// 每读取一行，就去统计该行有多少个英文、数字、空格和其他字符
	// 然后将结果存放到一个结构体中

	f, err := os.Open("./demo1.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// 创建reader
	reader := bufio.NewReader(f)

	// 创建结构体
	var charcount CharCount

	for {
		str, err := reader.ReadString('\n') //读取每一行
		if err == io.EOF {
			break
		}
		for _, value := range str {
			switch {
			case value >= 'a' && value <= 'z':
				fallthrough //	穿透
			case value >= 'A' && value <= 'Z':
				charcount.chCount++
			case value == ' ' || value == '\t':
				charcount.SpaceCount++
			case value >= '0' && value <= '9':
				charcount.NumCount++
			default:
				charcount.OtherCount++
			}
		}
	}
	fmt.Printf("英文个数：%v\n数字个数：%v\n空格个数：%v\n其他字符数字：%v\n",
		charcount.chCount, charcount.NumCount, charcount.SpaceCount, charcount.OtherCount)
}

// 接受命令行参数
func args() {
	// 定义几个变量，用于接受命令行的参数值
	var user string
	var pwd string
	var host string
	var port string

	// &user 就是接受用户命令行中输入的 -u 后面的参数值
	// "u" 就是 -u
	// "" 默认值
	// "用户名,默认为空" 是参数说明
	flag.StringVar(&user, "u", "", "用户名,默认为空")
	flag.StringVar(&pwd, "p", "", "密码,默认为空")
	flag.StringVar(&host, "h", "", "主机,默认为空")
	flag.StringVar(&port, "p", "", "端口,默认为空")

	// 命令行输入的是：
	// mysqld -u root -p 123456 -h 127.0.0.1 -p 3306
}
