package main

import (
	"fmt"
	"time"
)

func trace(funcName string) func() {
	fmt.Printf("%v 开始执行\n", funcName)
	start := time.Now()
	return func() {
		fmt.Printf("%v 执行结束, cost: %v\n", funcName, time.Since(start))
	}
}

func sum(x, y int) (res int) {
	defer func() {
		res++
		fmt.Println(res)
	}()
	return x + y
}

func after() {
	defer trace("111")()
}

func deferTest() {
	//defer trace("deferTest")()
	//time.Sleep(1 * time.Second)
	//sum(1, 1)
	for i := 0; i < 5; i++ {
		after()
		fmt.Println(222)
	}
}
