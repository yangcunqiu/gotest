package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

var lastClientAddr string
var coonMap = make(map[string]net.Conn)
var waitSendMessageMap = make(map[string][]string)
var coonAddrSlice = make([]string, 10)

func main() {
	// 启动服务器, 监听端口
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("服务器监听失败, err=%v\n", err)
		return
	}
	fmt.Println("服务器启动成功")
	defer listen.Close()

	// 一直等待客户端连接
	for {
		conn, err := listen.Accept()

		if err != nil {
			fmt.Printf("服务器获取连接失败, err=%v\n", err)
		}
		// 获取客户端地址
		clientAddr := conn.RemoteAddr().String()
		fmt.Printf("客户端[%v]连接成功\n", clientAddr)
		// 将当前客户端连接存入map
		coonMap[clientAddr] = conn
		coonAddrSlice = append(coonAddrSlice, clientAddr)
		// 处理客户端连接后的逻辑
		go process(conn)
	}
}

func broadcast() {
	for addr, coon := range coonMap {
		var message string
		for _, clientAddr := range coonAddrSlice {
			if clientAddr != "" && clientAddr != addr {
				message = message + clientAddr + "\n"
			}
		}

		if message != "" {
			message = "当前在线客户端地址: \n" + message
			err := send(coon, message)
			if err != nil {
				fmt.Printf("给客户端发送消息失败, err=%v", err)
			}
		}

	}
}

func process(conn net.Conn) {
	clientAddr := conn.RemoteAddr().String()
	// 广播当前在线客户端地址
	broadcast()

	// 发送留言
	sendWaitMessage(conn)

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Printf("客户端[%v]已退出\n", clientAddr)
			// 移除出map
			coonMap[clientAddr] = nil
			return
		} else if err != nil {
			fmt.Printf("接收客户端[%v]消息失败, err=%v\n", clientAddr, err)
		}

		clientMessage := string(buf[:n])
		heartMessage := clientAddr + "heart!!!" + clientAddr
		if clientMessage != heartMessage {
			if strings.Index(clientMessage, "$") != -1 {
				// 客户端给另一个客户端发消息
				sendClient(conn, clientMessage)
			} else {
				// 简单输出客户端消息
				if lastClientAddr == "" || lastClientAddr != clientAddr {
					fmt.Println(clientAddr + ":")
				}
				fmt.Print(clientMessage)
				lastClientAddr = clientAddr
			}
		}
	}
}

func sendWaitMessage(conn net.Conn) {
	clientAddr := conn.RemoteAddr().String()
	for addr, messageSlice := range waitSendMessageMap {
		if addr == clientAddr {
			var waitMessage string
			for _, message := range messageSlice {
				waitMessage += message
			}
			err := send(conn, waitMessage)
			if err != nil {
				fmt.Printf("给客户端发送消息失败, err=%v", err)
			}
		}
	}
	waitSendMessageMap[clientAddr] = nil
}

func sendClient(conn net.Conn, message string) {
	clientAddr := conn.RemoteAddr().String()
	// 解析出目的地址
	split := strings.Split(message, "|")
	destAddr := split[0][1:]
	destCoon := coonMap[destAddr]
	destMessage := clientAddr + ":\n" + split[1]
	if destCoon != nil {
		// 对方在线 直接发送
		err := send(destCoon, destMessage)
		if err != nil {
			fmt.Printf("客户端[%v] -> 客户端[%v]发送消息失败, err=%v", clientAddr, destAddr, err)
		}
		fmt.Printf("客户端[%v] -> 客户端[%v]发送消息, 消息内容=%v", clientAddr, destAddr, destMessage)
	} else {
		// 对方不在线 保存到等待发送的map
		waitSendMessage := destMessage
		waitSendMessageMap[destAddr] = make([]string, 10)
		waitSendMessageMap[destAddr] = append(waitSendMessageMap[destAddr], waitSendMessage)
	}
}

func send(conn net.Conn, message string) error {
	_, err := conn.Write([]byte(message))
	return err
}
