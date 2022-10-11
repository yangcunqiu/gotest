package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var pc = func() (pc [32]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()

func array() {
	// crypto/sha256包的Sum256函数对一个任意的字节slice类型的数据生成一个对应的消息摘要。
	// 消息摘要有256bit大小，因此对应[32]byte数组类型
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	// 练习 4.2： 编写一个程序，默认情况下打印标准输入的SHA256编码，并支持通过命令行flag定制，输出SHA384或SHA512哈希算法。
	practice42()
}

func practice42() {
	// os.args: -type 512
	fmt.Println("请输入任意字符串")
	var str string
	_, _ = fmt.Scan(&str)
	var a string
	flag.StringVar(&a, "type", "256", "获取加密算法类型")
	flag.Parse()
	switch a {
	case "256":
		fmt.Printf("SHA256编码为: %x", sha256.Sum256([]byte(str)))
	case "384":
		fmt.Printf("SHA384编码为: %x", sha512.Sum384([]byte(str)))
	case "512":
		fmt.Printf("SHA512编码为: %x", sha512.Sum512([]byte(str)))
	}
}
