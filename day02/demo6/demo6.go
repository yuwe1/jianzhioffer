package main

import "fmt"

// 计算二进制中1的个数(personal)
func Total_1(data int) int {
	if data == 0 {
		return 0
	}
	if data < 0 {
		data = -data
	}
	// 将数据转换成二进制
	sum := 0
	// 商
	m := data / 2
	// 余数
	n := data % 2
	// fmt.Println(n)
	for m != 0 {
		sum += n
		temp := m
		m = m / 2
		// 余数
		n = temp % 2
		// fmt.Println(n)
	}
	return sum + 1
}

// 递归求解
func Total_two(data int) int {
	if data == 0 {
		return 0
	}
	if data < 0 {
		data = -data
	}
	m := data / 2
	n := data % 2
	return n + Total_two(m)
}

// 思路四
func Total_three(data int) int {
	count := 0
	if data <= 0 {
		data = -data
	}
	for data != 0 {
		data = data & (data - 1)
		count++
	}
	return count
}
func main() {
	// fmt.Println(Total_1(15))
	// fmt.Println(Total_two(-15))
	// fmt.Println((-15) >> 1)
	fmt.Println(Total_three(-9))
	fmt.Println(Total_three(9))
	fmt.Println(9 & 8)
}
