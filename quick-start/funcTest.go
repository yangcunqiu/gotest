package main

import "fmt"

func add() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func funcTest() {

	//a := func() int {
	//	return 1
	//}()
	//fmt.Println(a)

	//f := func(x, y int) int {
	//	return x + y
	//}
	//
	//fmt.Println(f(1, 1))
	//fmt.Println(f(1, 1))
	//
	f := add()
	fmt.Printf("%v, %T\n", f(), f)
	fmt.Printf("%v, %T\n", f(), f)
	fmt.Printf("%v, %T\n", f(), f)

	f1 := add()
	fmt.Printf("%v, %T\n", f1(), f1)
	//fmt.Println(f()) // "1"
	//fmt.Println(f()) // "2"
	//fmt.Println(f()) // "3"

	//fmt.Println(add()) // 0xcbb0e0
	//fmt.Println(add()) // 0xcbb0c0
	//fmt.Println(add()) // 0xcbb0a0
	//
	//fmt.Println(add()()) // 1
	//fmt.Println(add()()) // 1
	//fmt.Println(add()()) // 1
}
