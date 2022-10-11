package main

import "fmt"

func remove1(slice []int, index int) []int {
	return slice[:index+copy(slice[index:], slice[index+1:])]
}

func remove2(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func rotate(slice []int) []int {
	i := 0
	j := len(slice) - 1
	for j >= i {
		slice[i], slice[j] = slice[j], slice[i]
		i++
		j--
	}

	return slice
}

func slice() {
	i := []int{1, 2, 3, 4, 5, 6}
	//fmt.Println(remove2(i, 2))
	fmt.Println(rotate(i))
}
