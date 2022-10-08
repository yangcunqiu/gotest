package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func dup1() {
	counts := make(map[string]int)
	// bufio包，它使处理输入和输出方便又高效。Scanner类型是该包最有用的特性之一，它读取输入并将其拆成行或单词；通常是处理行形式的输入最简单的方法
	input := bufio.NewScanner(os.Stdin)
	// input变量从程序的标准输入中读取内容。每次调用input.Scan()，即读入下一行，并移除行末的换行符；读取的内容可以调用input.Text()得到。
	// Scan函数在读到一行时返回true，不再有输入时返回false
	for input.Scan() {
		// counts[input.Text()]++语句等价下面两句：
		// line := input.Text()
		// counts[line] = counts[line] + 1
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			// os.Open函数返回两个值。第一个值是被打开的文件(*os.File），其后被Scanner读取
			// os.Open返回的第二个值是内置error类型的值
			f, err := os.Open(arg)
			// 如果err的值不是nil，说明打开文件时出错了
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			// 如果err等于内置值nil，那么文件被成功打开。读取文件，直到文件结束
			countLines(f, counts)
			// 然后调用Close关闭该文件
			_ = f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

func dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// ReadFile函数返回一个字节切片（byte slice），必须把它转换为string，才能用strings.Split分割
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func practice4() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			// os.Open函数返回两个值。第一个值是被打开的文件(*os.File），其后被Scanner读取
			// os.Open返回的第二个值是内置error类型的值
			f, err := os.Open(arg)
			// 如果err的值不是nil，说明打开文件时出错了
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			// 如果err等于内置值nil，那么文件被成功打开。读取文件，直到文件结束
			input := bufio.NewScanner(f)
			for input.Scan() {
				if counts[input.Text()] > 0 {
					fmt.Printf("重复行: %v, 重复行所在文件: %v\n", input.Text(), f.Name())
				}
				counts[input.Text()]++
			}
			// 然后调用Close关闭该文件
			_ = f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup() {
	// os.args: D:\go\gotest\quick-start\dupTest.txt
	// 对文件做拷贝、打印、搜索、排序、统计或类似事情的程序都有一个差不多的程序结构：一个处理输入的循环，在每个元素上执行计算处理，在处理的同时或最后产生输出
	// 灵感来自于Unix的uniq命令，其寻找相邻的重复行
	// dup1()

	// 很多程序要么从标准输入中读取数据，如上面的例子所示，要么从一系列具名文件中读取数据。
	// dup程序的下个版本读取标准输入或是使用os.Open打开各个具名文件，并操作它们
	// dup2()

	// dup的前两个版本以"流”模式读取输入，并根据需要拆分成多个行。理论上，这些程序可以处理任意数量的输入数据。还有另一个方法，就是一口气把全部输入数据读到内存中，一次分割为多行，然后处理它们
	// 这个例子引入了ReadFile函数（来自于io/ioutil包），其读取指定文件的全部内容，strings.Split函数把字符串分割成子串的切片。
	// dup3()

	// 练习 1.4： 修改dup2，出现重复的行时打印文件名称。
	practice4()
}
