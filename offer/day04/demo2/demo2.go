package main

import "fmt"

// 买卖股票的最大利润
func getProfi(data []int) int {
	if len(data) <= 1 {
		return 0
	}
	sum := 0
	for i := 1; i < len(data); i++ {
		if data[i] > data[i-1] {
			sum += (data[i] - data[i-1])
		}
	}
	return sum
}

func main() {
	data := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(getProfi(data))
}
