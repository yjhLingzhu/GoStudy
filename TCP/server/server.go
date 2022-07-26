package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		// 创建一个切片
		buf := make([]byte, 1024)
		fmt.Printf("服务器在等待%s客户端发信息！\n", conn.RemoteAddr().String())
		// 1. 等待客户端发送的信息
		// 2. 如果客户端没有write[发送], 那么协程会阻塞在这里直到超时
		n, err := conn.Read(buf) // 从conn读取
		if err != nil {
			fmt.Println("服务器端read err", err)
			break
		}
		// 3. 显示信息到服务器终端
		fmt.Print("信息是：", string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听....")
	// 监听tcp协议的8888端口
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err!", err)
		return
	}
	defer listen.Close()

	// 循环等待客户端来连接
	for {
		// 等待客户端连接，阻塞
		fmt.Println("等待客户端来连接...")
		conn, err := listen.Accept() // 阻塞
		if err != nil {
			fmt.Println("accept err!", err)
			continue
		} else {
			fmt.Printf("conn=%v, addr=%s\n", conn, conn.RemoteAddr().String())
		}
		// 这里开协程为客户端服务
		go handleConnection(conn)
	}
}
