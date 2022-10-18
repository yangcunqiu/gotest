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
	var p Person = User{"xm", 1}
	person, ok := p.(Person)
	fmt.Println(person, ok)

	athlete, ok := p.(Athlete)
	fmt.Println(athlete, ok)

	if per, ok := p.(Person); ok {
		fmt.Println(per, ok)
	}
}
