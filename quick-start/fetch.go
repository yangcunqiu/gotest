package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func fetch1() {
	for _, url := range os.Args[1:] {
		// http.Get函数是创建HTTP请求的函数，如果获取过程没有出错，那么会在resp这个结构体中得到访问的请求结果
		// resp的Body字段包括一个可读的服务器响应流
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// ioutil.ReadAll函数从response中读取到全部内容
		b, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close关闭resp的Body流，防止资源泄露
		_ = resp.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

func practice7() {
	for _, url := range os.Args[1:] {
		// http.Get函数是创建HTTP请求的函数，如果获取过程没有出错，那么会在resp这个结构体中得到访问的请求结果
		// resp的Body字段包括一个可读的服务器响应流
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		// resp.Body.Close关闭resp的Body流，防止资源泄露
		_ = resp.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		// fmt.Printf(os.Stdout)
	}
}

func practice8() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// ioutil.ReadAll函数从response中读取到全部内容
		b, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close关闭resp的Body流，防止资源泄露
		_ = resp.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

func practice9() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("http状态码: ", resp.Status)
		fmt.Printf("%s", b)
	}
}

func fetch() {
	// os.args: http://www.baidu.com
	// Go语言在net这个强大package的帮助下，使用这个包可以更简单地用网络收发信息，还可以建立更底层的网络连接，编写服务器程序。
	// 这个例子的灵感来源于curl工具（译注：unix下的一个用来发http请求的工具，具体可以man curl）
	// fetch1()

	// 练习 1.7： 函数调用io.Copy(dst, src)会从src中读取内容，并将读到的结果写入到dst中，
	// 使用这个函数替代掉例子中的ioutil.ReadAll来拷贝响应结构体到os.Stdout，避免申请一个缓冲区（例子中的b）来存储。记得处理io.Copy返回结果中的错误。
	// practice7()

	// 练习 1.8： 修改fetch这个范例，如果输入的url参数没有 http:// 前缀的话，为这个url加上该前缀。你可能会用到strings.HasPrefix这个函数。
	// os.args: www.baidu.com
	// practice8()

	// 练习 1.9： 修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码。
	// os.args: http://www.baidu.com
	practice9()
}
