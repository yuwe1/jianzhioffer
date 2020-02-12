package main

import "fmt"

// 数字第一次出现的位置
func getFirstName(data []int, target int) int {
	if len(data) == 0 {
		return -1
	}
	i, j := 0, len(data)-1

	for i <= j {
		mid := (i + j) >> 1
		if target == data[mid] {
			return mid
		} else if target < data[mid] {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return -1
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(getFirstName(data, 0))
}
