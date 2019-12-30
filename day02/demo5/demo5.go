package main

import "fmt"

var (
	data map[rune]int
)

func init() {
	data = make(map[rune]int, 26)
	for i := 'A'; i <= 'Z'; i++ {
		data[i] = int(i) - 64
	}
}
func TwentysixToTen(str string) int {
	if len(str) <= 0 {
		return 0
	}
	sum := 0
	for i, _ := range str {
		j := len(str) - 1 - i
		m := Calaute(i)
		sum += (data[rune(str[j])] * m)
	}
	return sum
}

// 计算26的几次
func Calaute(n int) int {
	sum := 1
	for i := 0; i < n; i++ {
		sum *= 26
	}
	return sum
}

// 第二种解法降低空间复杂度O(1)
func TwentysixToTen_two(str string) int {
	sum := 0
	for i := len(str) - 1; i >= 0; i-- {
		// 确定了每一个字符所对应的整数
		p := int(str[i])
		if (p-64) > 26 || (p-64) < 1 {
			return 0
		}
		m := Calaute(len(str) - 1 - i)
		sum += (p - 64) * m
	}
	return sum
}
func main() {
	// fmt.Println(TwentysixToTen("A"))
	fmt.Println(TwentysixToTen("AB"))
	fmt.Println(TwentysixToTen_two("Ab"))
}
