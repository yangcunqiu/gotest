package main

import "fmt"

type Person interface {
	eat(food string)
	walk() int
}

type Athlete interface {
	Person
	running() int
}

type User struct {
	name string
	age  int
}

func (u User) eat(food string) {
	fmt.Println(food)
}

func (u User) walk() int {
	return 100
}

func test(p Person) {
	fmt.Println(p == nil)
	if p != nil {
		fmt.Println(p.walk())
	}
}

func interfaceTest() {
	//var u *User
	var u Person
	fmt.Println(u == nil)
	test(u)
}
