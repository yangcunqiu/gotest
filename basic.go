package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func numberCombinations() {
	// 有 1、2、3、4 这四个数字，能组成多少个互不相同且无重复数字的三位数？都是多少？
	count := 0
	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			for k := 1; k < 5; k++ {
				if i != j && i != k && j != k {
					count++
					fmt.Printf("%v\n", i*100+j*10+k)
				}
			}
		}
	}
	fmt.Printf("count: %v", count)
}

func calculateProfit() {
	/*
		企业发放的奖金根据利润提成。利润(I)低于或等于 10 万元时，奖金可提成 10%；利润高于 10 万元，低于 20 万元，低于 10 万元的部分按 10% 提成，高于 10 万元的部分，可提成 7.5%。
		20 万到 40 万之间时，高于 20 万元的部分，可提成 5%；40 万到 60 万之间时高于 40 万元的部分，可提成 3%；60 万到 100 万之间时，高于 60 万元的部分，可提成 1.5%，高于 100 万元时，超过 100 万元的部分按 1% 提成。
		从键盘输入当月利润 I，求应发放奖金总数？
	*/
	fmt.Printf("请输入当月利润: ")
	var profit float64
	_, _ = fmt.Scanln(&profit)
	var bonus float64
	switch {
	case profit > 1000000:
		bonus += (profit - 1000000) * 0.01
		profit = 1000000
		fallthrough
	case profit > 600000:
		bonus += (profit - 600000) * 0.015
		profit = 600000
		fallthrough
	case profit > 400000:
		bonus += (profit - 400000) * 0.03
		profit = 400000
		fallthrough
	case profit > 200000:
		bonus += (profit - 200000) * 0.05
		profit = 200000
		fallthrough
	case profit > 100000:
		bonus += (profit - 100000) * 0.075
		profit = 100000
		fallthrough
	default:
		bonus += profit * 0.1
	}
	fmt.Printf("bonus: %v\n", bonus)
}

func square() {
	// 完全平方指用一个整数乘以自己例如1*1，2*2，3*3等，依此类推。若一个数能表示成某个整数的平方的形式，则称这个数为完全平方数
	// 一个整数，它加上 100 后是一个完全平方数，再加上 168 又是一个完全平方数，请问该数是多少？
	num := 0
	for {
		sqrt1 := int(math.Sqrt(float64(num + 100)))
		sqrt2 := int(math.Sqrt(float64(num + 100 + 168)))
		if sqrt1*sqrt1 == num+100 && sqrt2*sqrt2 == num+100+168 {
			break
		}
		num++
	}
	fmt.Println("改数是: ", num)
}

func judgeDate() {
	// 输入某年某月某日，判断这一天是这一年的第几天？
	fmt.Println("请输入日期, (某年某月某日)")
	var datetimeStr string
	_, _ = fmt.Scan(&datetimeStr)
	yearSplit := strings.Split(datetimeStr, "年")
	yearStr := yearSplit[0]
	var year, _ = strconv.ParseInt(yearStr, 0, 0)
	monthSplit := strings.Split(yearSplit[1], "月")
	monthStr := monthSplit[0]
	var month, _ = strconv.ParseInt(monthStr, 0, 0)
	daySplit := strings.Split(monthSplit[1], "日")
	dayStr := daySplit[0]
	var day, _ = strconv.ParseInt(dayStr, 0, 0)
	dayCount := day
	switch {
	case month > 11:
		dayCount += 30
		fallthrough
	case month > 10:
		dayCount += 31
		fallthrough
	case month > 9:
		dayCount += 30
		fallthrough
	case month > 8:
		dayCount += 31
		fallthrough
	case month > 7:
		dayCount += 31
		fallthrough
	case month > 6:
		dayCount += 30
		fallthrough
	case month > 5:
		dayCount += 31
		fallthrough
	case month > 4:
		dayCount += 30
		fallthrough
	case month > 3:
		dayCount += 31
		fallthrough
	case month > 2:
		if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
			dayCount += 29
		} else {
			dayCount += 28
		}
		fallthrough
	case month > 1:
		dayCount += 31
	}
	fmt.Printf("这一天是这一年的第%d天", dayCount)
}

func comparisonSize() {
	// 输入三个 整数 x，y，z，请把这三个数由小到大输出
	fmt.Println("请输入3个数, 空格隔开")
	var x int
	var y int
	var z int
	_, _ = fmt.Scanf("%d %d %d", &x, &y, &z)
	if x > y {
		x, y = y, x
	}
	if x > z {
		x, z = z, x
	}
	if y > z {
		y, z = z, y
	}
	fmt.Println(x, y, z)
}

func printMultiplicationTable() {
	// 输出 9*9 乘法口诀表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d  ", i, j, i*j)
		}
		println()
	}
}

func judgePrimeNumber() {
	// 质数又称素数。一个大于1的自然数，除了1和它自身外，不能被其他自然数整除的数叫做质数；否则称为合数（规定1既不是质数也不是合数）。
	// 判断 101-200 之间有多少个素数，并输出所有素数
	var count int
	for i := 101; i < 201; i++ {
		num := int(math.Sqrt(float64(i)))
		isMatch := false
		for j := 2; j <= num; j++ {
			if i%j == 0 && j != i {
				isMatch = true
				break
			}
		}
		if !isMatch {
			println(i)
			count++
		}
	}
	fmt.Printf("共有%d个素数", count)
}

func daffodilNumber() {
	// 打印出所有的 “水仙花数”，所谓 “水仙花数” 是指一个三位数，其各位数字立方和等于该数本身。例如：153 是一个 “水仙花数”，因为 153=1 的三次方＋5 的三次方＋3 的三次方
	for num := 100; num < 1000; num++ {
		i := num / 100
		j := num / 10 % 10
		k := num % 10
		if i*i*i+j*j*j+k*k*k == num {
			println(num)
		}
	}
}

func statisticalStringCharCount() {
	// 用 Golang 实现，统计一个 字符串 中各个字符的个数
	fmt.Printf("请输入一个字符串, 可包含英文字母、空格、数字和其它字符")
	var str string
	i, j, k, l, m := 0, 0, 0, 0, 0
	_, _ = fmt.Scan(&str)
	for _, ch := range str {
		switch {
		case ch >= 'A' && ch <= 'Z':
			i++
		case ch >= 'a' && ch <= 'z':
			j++
		case ch == ' ':
			k++
		case ch >= '0' && ch <= '9':
			l++
		default:
			m++
		}
	}
	fmt.Printf("共%d个字符, 小写字母%d个, 大写字母%d个, 空格%d个, 数字%d个, 其他字符%d个", i+j+k+l+m, i, j, k, l, m)
}

func main() {
	// numberCombinations()
	// calculateProfit()
	// square()
	// judgeDate()
	// comparisonSize()
	// printMultiplicationTable()
	// judgePrimeNumber()
	// daffodilNumber()
	statisticalStringCharCount()
}
