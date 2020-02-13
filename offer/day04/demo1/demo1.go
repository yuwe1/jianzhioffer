package main

import "fmt"

// 数组实现加1
func add(nums []int) []int {
	if len(nums) == 0 {
		return []int{1}
	}
	l := len(nums)
	for i := l - 1; i >= 0; i-- {
		if nums[i] < 9 {
			nums[i] += 1
			return nums
		} else {
			nums[i] = 0
		}
	}
	return append([]int{1}, nums...)
}

func main() {
	data := []int{1, 2, 4}
	data2 := []int{9, 9, 9}
	data3 := []int{1, 9, 9}
	fmt.Println(add(data))
	fmt.Println(add(data3))
	fmt.Println(add(data2))
}
