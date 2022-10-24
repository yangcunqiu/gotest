package main

import "fmt"

func mapTest() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	for key := range ages {
		fmt.Println(key)
	}

	for key, value := range ages {
		fmt.Println(key, value)
	}

	for _, value := range ages {
		fmt.Println(value)
	}
}
