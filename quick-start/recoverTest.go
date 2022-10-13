package main

import (
	"errors"
	"fmt"
)

func getInfo(id int) (err error) {
	defer func() {
		if p := recover(); p != nil {
			switch p {
			case "0":
				err = errors.New("id not found error")
			default:
				panic(p)
			}
		}
	}()

	if id == 0 {
		panic("0")
	} else if id == 1 {
		panic("121212")
	}
	fmt.Println(id)
	return nil
}

func recoverTest() {
	//err1 := getInfo(0)
	//if err1 != nil {
	//	fmt.Println(err1)
	//}
	//err2 := getInfo(1)
	//if err2 != nil {
	//	fmt.Println(err2)
	//}
	err := getInfo(2)
	if err != nil {
		fmt.Println(err)
	}
}
