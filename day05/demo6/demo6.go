package main

import "fmt"

func duplicateInArray(nums []int) int {
	// 定义一个map
	m := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return nums[i]
		} else {
			m[nums[i]] = true
		}
	}
	return -1
}

func main() {
	data := []int{1, 7, 5, 9, 7, 9, 5, 1, 5, 6, 7}
	fmt.Println(duplicateInArray(data))
}
