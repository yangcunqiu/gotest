package main

import (
	"fmt"
	"time"
)

func channelTest1() {
	intChan := make(chan int, 5)
	intChan <- 1
	intChan <- 2
	intChan <- 3
	fmt.Printf("%v %T len=%v cap=%v\n", intChan, intChan, len(intChan), cap(intChan))

	var int1 int
	int1 = <-intChan
	int2 := <-intChan
	fmt.Println(int1, int2)
	fmt.Printf("%v %T len=%v cap=%v\n", intChan, intChan, len(intChan), cap(intChan))
	<-intChan
	fmt.Printf("%v %T len=%v cap=%v\n", intChan, intChan, len(intChan), cap(intChan))
	int3 := <-intChan
	fmt.Println(int3)
}

func channelTest2() {
	intChan := make(chan int, 5)
	intChan <- 1
	intChan <- 2
	intChan <- 3

	close(intChan)
	fmt.Println(<-intChan)
	fmt.Println(<-intChan)
	fmt.Println(<-intChan)
	fmt.Println(<-intChan)

	//intChan <- 4

	close(intChan)
	//var a chan int
	//close(a)
}

func channelTest3() {
	intChan1 := make(chan int)
	intChan2 := make(chan int, 0)
	intChan3 := make(chan int, 1)
	fmt.Printf("intChan1 cap = %v, intChan2 cap = %v, intChan3 cap = %v\n", cap(intChan1), cap(intChan2), cap(intChan3))
}

func channelTest4() {
	done := make(chan struct{})
	go func() {
		fmt.Println("func done!")
		done <- struct{}{}
	}()
	<-done
	fmt.Println("main done!")
}

func getNaturals(in chan<- int) {
	for i := 0; i < 100; i++ {
		in <- i
	}
	close(in)
}

func getSquares(out <-chan int, in chan<- int) {
	for num := range out {
		in <- num * num
	}
	close(in)
}

func getResult(out <-chan int) {
	for res := range out {
		fmt.Println(res)
	}
}

func channelTest5() {
	naturals := make(chan int)
	squares := make(chan int)

	go getNaturals(naturals)
	go getSquares(naturals, squares)
	go getResult(squares)

	time.Sleep(10 * time.Second)
}

func channelTest6() {
	intChan := make(chan int, 3)
	intChan <- 1
	intChan <- 2
	intChan <- 3
	intChan <- 4
	fmt.Println(len(intChan), cap(intChan))
	fmt.Println(<-intChan)
	fmt.Println(<-intChan)
	fmt.Println(<-intChan)
	fmt.Println(<-intChan)
}

func channelTest() {
	//channelTest1()
	//channelTest2()
	//channelTest3()
	//channelTest4()
	//channelTest5()
	channelTest6()
}
