package main

import (
	"fmt"
)

func ReOrderOddEven(data []int) []int {
	if len(data) == 0{
		return data
	}
	var temp1 []int
	var temp2 []int
	lengh := len(data)
	for i := 0; i < lengh; i++ {
		if data[i]&1 == 1 {
			temp1 = append(temp1, data[i])
		} else {
			temp2 = append(temp2, data[i])
		}
	}
	var result []int
	for i := 0; i < len(temp1); i++ {
		result = append(result, temp1[i])
	}
	for i := 0; i < len(temp2); i++ {
		result = append(result, temp2[i])
	}
	return result
}

// 思路二
func ReOrderOddEven_two(data []int) []int {
	// 定义头尾指针
	i, j := 0, len(data)-1
	for i <= j {
		if data[i]&1 == 0 && data[j]&1 == 1 {
			data[i], data[j] = data[j], data[i]
			i++
			j--
		} else if data[i]&1 == 1 && data[j]&1 == 1 {
			i++
		} else if data[i]&1 == 0 && data[j]&1 == 0 {
			j--
		} else {
			i++
			j--
		}

	}
	return data
}

// 基于思路二的扩展思路三
func ReOrderOddEven_three(data []int) []int {
	if len(data) <= 0 {
		return data
	}
	i, j := 0, len(data)-1
	for i <= j {
		for data[i]&1 == 1 && i <= j {
			i++
		}
		for data[j]&1 == 0 && i <= j {
			j--
		}
		if i <= j {
			data[i], data[j] = data[j], data[i]
		}

	}
	return data
}

// 更好的思路，具有模板性
type F func(int) bool

// {0, 1, 2, 7, 9, 1, 0}
func ReOrderOddEven_four(data []int, f F) []int {
	if len(data) <= 0 {
		return data
	}
	i, j := 0, len(data)-1
	for i <= j {
		for f(data[i]) && i <= j {
			i++
		}
		for !f(data[j]) && i <= j {
			j--
		}
		if i <= j {
			fmt.Println("交换")
			data[i], data[j] = data[j], data[i]
		}

	}
	return data
}

// 奇数偶数判断，奇数返回true
func JudgeOddAndEven(data int) bool {
	if data&1 == 1 {
		return true
	}
	return false
}

// 能被3整除的放左边，其余放右边
func JudgeDivide3(data int) bool {
	if data%3 == 0 && data >= 3 {
		return true
	}
	return false
}

// 主函数
func main() {
	data := []int{0, 1, 2, 7, 9, 1, 0}
	// fmt.Println(ReOrderOddEven(data))
	// data = []int{0, 1, 2, 7, 9, 1, 0}
	// fmt.Println(ReOrderOddEven_two(data))
	// data = []int{0, 1, 2, 7, 9, 1, 0}
	// fmt.Println(ReOrderOddEven_three(data))
	// data = []int{0, 1, 2, 7, 9, 1, 0}
	fmt.Println(ReOrderOddEven_four(data, JudgeOddAndEven))
	// data = []int{0, 1, 2, 7, 9, 1, 0}
	// fmt.Println(ReOrderOddEven_four(data, JudgeDivide3))
}
