package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func concurrencyTest1(intChan chan int) {
	var wg sync.WaitGroup
	for num := range intChan {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println(num)
		}(num)
	}
	wg.Wait()
}

func selectTest() {
	abort := make(chan struct{})
	go func() {
		_, _ = os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	tick := time.Tick(1 * time.Second)
	intChan := make(chan struct{})
	go func() {
		for i := 10; i >= 0; i-- {
			fmt.Println(i)
			<-tick
		}
		intChan <- struct{}{}
	}()

	select {
	case <-abort:
		fmt.Println("手动停止!!!")
		return
	case <-intChan:
		fmt.Println("火箭发射")
	}
}

func selectTest2() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x) // "0" "2" "4" "6" "8"
		case ch <- i:
		}
	}
}

func selectTest3() {
	ch := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		ch <- struct{}{}
	}()

	for {
		select {
		case <-ch:
			fmt.Println("ch-case就绪")
			return
		default:
			fmt.Println("没有任何case就绪")
		}
	}
}

func selectTest4() {
	for i := 0; i < 10; i++ {
		go func(num int) {
			fmt.Println("go开始执行", num)
		}(i)
	}

	time.Sleep(5 * time.Second)
}

func concurrencyTest() {
	intChan := make(chan int, 3)
	intChan <- 1
	intChan <- 2
	intChan <- 3
	close(intChan)

	//concurrencyTest1(intChan)
	//selectTest()
	//selectTest2()
	//selectTest3()
	selectTest4()
}
