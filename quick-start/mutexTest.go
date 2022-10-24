package main

import (
	"fmt"
	"sync"
)

var (
	mu  sync.Mutex
	rw  sync.RWMutex
	num int
)

func mutexTest1() {
	wg := sync.WaitGroup{}
	wg.Add(1100)
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			num++
			wg.Done()
		}()
	}
	for i := 0; i < 100; i++ {
		go func() {
			rw.Lock()
			defer rw.Unlock()
			fmt.Println(num)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(num)
}

func mutexTest2() {
	var x, y int
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		x = 1 // A1
		fmt.Print("y:", y, " ")
		wg.Done()
	}()
	go func() {
		y = 1
		fmt.Print("x:", x, " ")
		wg.Done()
	}()
	wg.Wait()
}

func mutexTest() {
	//mutexTest1()
	mutexTest2()
}
