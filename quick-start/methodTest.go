package main

import "fmt"

func (p Point) getSum(num int) int {
	return p.X + p.Y + num
}

func (p Point) getSub(num int) int {
	return p.X - p.Y - num
}

func (p Point) updateX() {
	p.X = 2
}

//func (p *Point) updatePtrX() {
//	if p != nil {
//		p.X = 2
//	}
//}

func methodTest() {
	p1 := Point{X: 1, Y: 1}
	//fmt.Println(*p1)
	//p1.updateX()
	//p1.updatePtrX()
	//m := p1.updatePtrX
	//fmt.Printf("%T", m)
	//m()
	//fmt.Println(*p1)

	m1 := Point.getSub
	m2 := Point.getSum
	fmt.Printf("%T\n", m1)
	fmt.Println(m1(p1, 1))
	fmt.Println(m2(p1, 1))

}
