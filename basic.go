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

func calculateProfit1() {
	/*
		企业发放的奖金根据利润提成。利润(I)低于或等于 10 万元时，奖金可提成 10%；利润高于 10 万元，低于 20 万元，低于 10 万元的部分按 10% 提成，高于 10 万元的部分，可提成 7.5%。
		20 万到 40 万之间时，高于 20 万元的部分，可提成 5%；40 万到 60 万之间时高于 40 万元的部分，可提成 3%；60 万到 100 万之间时，高于 60 万元的部分，可提成 1.5%，高于 100 万元时，超过 100 万元的部分按 1% 提成。
		从键盘输入当月利润 I，求应发放奖金总数？
	*/
	fmt.Printf("请输入当月利润: ")
	var profit float64
	fmt.Scanln(&profit)
	var bonus float64
	if profit <= 100000 {
		bonus = profit * 0.1
	} else if profit > 100000 && profit <= 200000 {
		bonus = 100000*0.1 + (profit-100000)*0.075
	} else if profit > 200000 && profit <= 400000 {
		bonus = 100000*0.1 + (profit-100000)*0.075 + (profit-200000)*0.05
	} else if profit > 400000 && profit <= 600000 {
		bonus = 100000*0.1 + (profit-100000)*0.075 + (profit-200000)*0.05 + (profit-400000)*0.03
	} else if profit > 600000 && profit <= 1000000 {
		bonus = 100000*0.1 + (profit-100000)*0.075 + (profit-200000)*0.05 + (profit-400000)*0.03 + (profit-600000)*0.015
	} else if profit > 1000000 {
		bonus = 100000*0.1 + (profit-100000)*0.075 + (profit-200000)*0.05 + (profit-400000)*0.03 + (profit-600000)*0.015 + (profit-1000000)*0.01
	}
	fmt.Printf("bonus: %v\n", bonus)
}

func calculateProfit2() {
	fmt.Printf("请输入当月利润: ")
	var profit float64
	fmt.Scanln(&profit)
	var bonus float64
	switch {

	}
	fmt.Printf("bonus: %v/n", bonus)
}

func main() {
	numberCombinations()
	calculateProfit1()
	calculateProfit2()
}
