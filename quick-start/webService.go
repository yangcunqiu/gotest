package main

import (
	"fmt"
	"log"
	"net/http"
)

func webService() {
	// 展示一个微型服务器，这个服务器的功能是返回当前用户正在访问的URL
	webService1()
}

func webService1() {
	// 将所有发送到/路径下的请求和handler函数关联起来
	http.HandleFunc("/", handler)
	// /开头的请求其实就是所有发送到当前站点上的请求，服务监听8000端口
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
