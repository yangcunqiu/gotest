package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func getUrl() (string, error) {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(5)%2 == 0 {
		return "", errors.New("errorMessage")
	}
	return "success", nil
}

func errorTest() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		url, err := getUrl()
		if err == nil {
			fmt.Println(url)
			break
		}
	}
}
