package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func practice1() {
	for _, arg := range os.Args {
		fmt.Println("1.1--" + arg)
	}
}

func practice2() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("1.2--[%v][%v]\n", i, arg)
	}
}

func practice3() {
	tempSlice := make([]string, 100000)
	for i := 0; i < 100000; i++ {
		tempSlice[i] = "temp-" + string(i)
	}
	startTime := time.Now().Unix()
	fmt.Println(strings.Join(tempSlice, " "))
	endTime := time.Now().Unix()
	highTime := endTime - startTime
	fmt.Printf("cost %d ms\n", highTime)

	startTimeLow := time.Now().Unix()
	for _, arg := range tempSlice {
		fmt.Println("1.3--" + arg)
	}
	endTimeLow := time.Now().Unix()
	lowTime := endTimeLow - startTimeLow
	fmt.Printf("low cost %d ms\n", lowTime)
	fmt.Printf("frugal time : %d ms", lowTime-highTime)
}

func echo() {
	// os包以跨平台的方式，提供了一些与操作系统交互的函数和变量。程序的命令行参数可从os包的Args变量获取；os包外部使用os.Args访问该变量。
	// 下面是Unix里echo命令的一份实现，echo把它的命令行参数打印成一行
	var s1, sep string
	for i := 1; i < len(os.Args); i++ {
		s1 += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s1)
	// for循环的另一种形式，在某种数据类型的区间（range）上遍历，如字符串或切片。echo的第二版本展示了这种形式：
	s2, sep := "", ""
	for _, arg := range os.Args[1:] {
		s2 += sep + arg
		sep = " "
	}
	fmt.Println(s2)
	// 上面这两种如果连接涉及的数据量很大，这种方式代价高昂。一种简单且高效的解决方案是使用strings包的Join函数
	fmt.Println(strings.Join(os.Args[1:], " "))
	// 练习 1.1： 修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。
	practice1()
	// 练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行
	practice2()
	// 练习 1.3： 做实验测量潜在低效的版本和使用了strings.Join的版本的运行时间差异。（1.6节讲解了部分time包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测
	practice3()
}
