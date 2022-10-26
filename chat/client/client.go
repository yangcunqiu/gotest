package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("连接服务器失败, err=%v", err)
		return
	}
	// 关闭连接
	defer conn.Close()
	fmt.Printf("连接服务器[%v]成功\n", conn.RemoteAddr().String())

	// 开启心跳检测
	go heart(conn)

	go func() {
		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err == nil {
				fmt.Println(string(buf[:n]))
			}
		}
	}()

	// 向服务器发消息
	for {
		// 接收控制台输入
		input := getInput()
		if input == "exit\r\n" {
			fmt.Println("已退出")
			return
		}
		// 向服务器发消息
		err = send(conn, input)
		if err != nil {
			fmt.Printf("向服务器发送消息失败, err=%v", err)
		}
	}

}

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	readString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("输入错误, err=%v\n", err)
	}
	return readString
}

func send(conn net.Conn, message string) error {
	_, err := conn.Write([]byte(message))
	return err
}

func heart(conn net.Conn) {
	tick := time.Tick(time.Second)
	for {
		// 每隔1s向服务器发送心跳检测, 如果发送失败了说明服务器已经关闭了, 需要退出客户端
		localAddr := conn.LocalAddr().String()
		err := send(conn, localAddr+"heart!!!"+localAddr)
		if err != nil {
			// 退出当前线程
			fmt.Println("服务器已关闭, 自动退出1")
			os.Exit(-1)
		}
		<-tick
	}
}
