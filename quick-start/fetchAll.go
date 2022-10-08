package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetchAll1() {
	start := time.Now()
	// make函数创建了一个传递string类型参数的channel
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// go function则表示创建一个新的goroutine，并在这个新的goroutine中执行这个函数
		go fetchUrl(url, ch) // start a goroutine
	}
	// 由main函数里的第二个for循环来处理并打印channel里的这个字符串。
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchAll() {
	// os.args: https://golang.org http://gopl.io https://godoc.org
	// 体验一下Go语言里的goroutine和channel
	// goroutine是一种函数的并发执行方式，而channel是用来在goroutine之间进行参数传递。main函数本身也运行在一个goroutine中
	// 面的例子fetchAll，和前面小节的fetch程序所要做的工作基本一致，fetchAll的特别之处在于它会同时去获取所有的URL，
	// 所以这个程序的总执行时间不会超过执行时间最长的那一个任务，前面的fetch程序执行时间则是所有任务执行时间之和
	fetchAll1()

	// os.args: https://www.runoob.com/go/go-slice.html https://www.runoob.com/go/go-program-structure.html http://c.biancheng.net/view/27.html
	// 练习 1.10： 找一个数据量比较大的网站，用本小节中的程序调研网站的缓存策略，
	// 对每个URL执行两遍请求，查看两次时间是否有较大的差别，并且每次获取到的响应内容是否一致，修改本节中的程序，将响应结果输出，以便于进行对比。
	practice10()

	// 练习 1.11： 在fetchall中尝试使用长一些的参数列表，比如使用在alexa.com的上百万网站里排名靠前的。
	// 如果一个网站没有回应，程序将采取怎样的行为？（Section8.9 描述了在这种情况下的应对机制）。
}

func practice10() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchUrl1(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchUrl1(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	// io.Copy会把响应的Body内容拷贝到ioutil.Discard输出流中
	// 可以把这个变量看作一个垃圾桶，可以向里面写一些不需要的数据

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	// 每当请求返回内容时，fetch函数都会往ch这个channel里写入一个字符串
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

func fetchUrl(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	// io.Copy会把响应的Body内容拷贝到ioutil.Discard输出流中
	// 可以把这个变量看作一个垃圾桶，可以向里面写一些不需要的数据
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	// 每当请求返回内容时，fetch函数都会往ch这个channel里写入一个字符串
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
