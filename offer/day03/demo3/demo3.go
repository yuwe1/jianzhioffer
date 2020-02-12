package main

import "fmt"

// 无序数组找出中位数
func FindMedian(data []int) {
	low := 0
	high := len(data) - 1
	mid := (high) / 2
	div := helper_two(data, low, high)
	for div != mid {
		if mid < div {
			div = helper_two(data, low, div-1)
		} else {
			div = helper_two(data, div+1, high)
		}
	}
	fmt.Println(data[mid])
}

//
func helper_two(data []int, low int, high int) int {
	left := low
	right := high
	// 关键字
	mid := right
	for left <= right {
		for left <= mid && data[left] <= data[mid] {
			left++
		}
		if left <= mid {
			data[left], data[mid] = data[mid], data[left]
			mid = left
		}
		for right >= mid && data[right] >= data[mid] {
			right--
		}
		if right >= mid {
			data[right], data[mid] = data[mid], data[right]
			mid = right
		}
	}
	return mid
}
func main() {
	data := []int{10, 6, 1, 5, 7}
	FindMedian(data)
}
