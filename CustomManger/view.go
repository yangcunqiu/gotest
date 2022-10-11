package main

import "fmt"

var indexCount = 0

func main() {
	for {
		fmt.Println("----------客户信息管理系统软件-----------")
		fmt.Println()
		fmt.Println("            1 添加客户")
		fmt.Println("            2 修改客户")
		fmt.Println("            3 删除客户")
		fmt.Println("            4 客户列表")
		fmt.Println("            5 退出")
		fmt.Println()
		var op int
		fmt.Println("            请选择(1-5):")
		_, _ = fmt.Scan(&op)
		switch op {
		case 1:
			fmt.Println("--------添加用户--------")
			var (
				name   string
				gender string
				age    int
				tel    string
				email  string
			)
			fmt.Print("姓名: ")
			_, _ = fmt.Scan(&name)
			fmt.Print("性别: ")
			_, _ = fmt.Scan(&gender)
			fmt.Print("年龄: ")
			_, _ = fmt.Scan(&age)
			fmt.Print("电话: ")
			_, _ = fmt.Scan(&tel)
			fmt.Print("邮箱: ")
			_, _ = fmt.Scan(&email)
			cus := custom{indexCount, name, gender, age, tel, email, false}
			add(cus)
			indexCount++
			fmt.Println("--------添加完成--------")
		case 2:
			fmt.Println("--------开始修改--------")
			fmt.Println("请输入想要修改的编号")
			var index int
			_, _ = fmt.Scan(&index)
			cus := get(index)
			fmt.Println(cus)
			var (
				name   string = cus.name
				gender string = cus.gender
				age    int    = cus.age
				tel    string = cus.tel
				email  string = cus.email
			)
			fmt.Print("姓名: ")
			_, _ = fmt.Scan(&name)
			fmt.Print("性别: ")
			_, _ = fmt.Scan(&gender)
			fmt.Print("年龄: ")
			_, _ = fmt.Scan(&age)
			fmt.Print("电话: ")
			_, _ = fmt.Scan(&tel)
			fmt.Print("邮箱: ")
			_, _ = fmt.Scan(&email)
			cus = custom{index, name, gender, age, tel, email, false}
			update(cus)
			fmt.Println("--------修改完成--------")
		case 3:
			fmt.Println("--------开始删除--------")
			fmt.Println("请输入要删除的编号")
			var id int
			_, _ = fmt.Scan(&id)
			del(id)
			fmt.Println("--------删除完成--------")
		case 4:
			fmt.Println("--------开始查询--------")
			for _, cus := range list() {
				fmt.Println(cus)
			}
			fmt.Println("--------查询完成--------")
		case 5:
			fmt.Println("已退出")
			goto EXIT
		}
	}
EXIT:
}
