package main

import (
	"encoding/json"
	"fmt"
)

type Hobby struct {
	HobbyName string
}

//type User struct {
//	Name string
//	Age  int
//	Hobby
//}

func jsonTest() {
	//user := User{Name: "ycq", Age: 0}
	//
	//data1, err := json.Marshal(user)
	//if err != nil {
	//	fmt.Printf("序列化失败: %v", err)
	//}
	//
	//data2, _ := json.MarshalIndent(user, "", "    ")
	//if err != nil {
	//	fmt.Printf("序列化失败: %v", err)
	//}
	//fmt.Println(string(data1))
	//fmt.Println(string(data2))

	var user2 User
	//str := "{\"name\":\"ycq\", \"age\":22, \"HobbyName\":\"足球\"}"
	str := `
			{"name":"ycq", "age":22, "hobby":{"hobbyName":"足球"}}
	`
	_ = json.Unmarshal([]byte(str), &user2)
	fmt.Println(user2)
}
