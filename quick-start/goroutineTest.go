package main

import (
	"fmt"
	"runtime"
	"time"
)

func goroutineTest() {
	//go spinner(100 * time.Millisecond)
	//const n = 45
	//fibN := fib(n) // slow
	//fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	fmt.Println(runtime.NumCPU())
	go test1()
	for i := 0; i < 100; i++ {
		fmt.Println("main", i)
		time.Sleep(1 * time.Second)
	}
}

func test1() {
	for i := 0; i < 100; i++ {
		fmt.Println("test1", i)
		time.Sleep(1 * time.Second)
	}
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
