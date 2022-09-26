package main

import "fmt"

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
	fmt.Scanln(&profit)
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
	// 一个整数，它加上 100 后是一个完全平方数，再加上 168 又是一个完全平方数，请问该数是多少？

}

func main() {
	//numberCombinations()
	//calculateProfit()
}
