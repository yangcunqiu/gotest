package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
	X      int
}

func structTest() {
	var s1 struct{}
	s2 := struct{}{}
	var p1 Point
	p2 := Point{}
	p3 := Point{1, 2}
	p4 := Point{Y: 2}
	fmt.Printf("%v, %T\n", s1, s1)
	fmt.Printf("%v, %T\n", s2, s2)
	fmt.Printf("%v, %T\n", p1, p1)
	fmt.Printf("%v, %T\n", p2, p2)
	fmt.Printf("%v, %T\n", p3, p3)
	fmt.Printf("%v, %T\n", p4, p4)

	c1 := Circle{
		Point:  Point{X: 1, Y: 8},
		Radius: 5,
		X:      10,
	}
	c2 := Circle{Point{1, 8}, 5, 10}
	fmt.Println(c1)
	fmt.Println(c2)
	var c Circle
	c.X = 1
	fmt.Println(c)
}
