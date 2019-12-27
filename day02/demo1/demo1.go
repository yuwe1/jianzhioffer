package main

import "fmt"

// 二分查找
func BirarySelect(data []int, key int) bool {
	low := 0
	high := len(data) - 1
	for low <= high {
		if key == data[(low+high)/2] {
			return true
		} else if key > data[(low+high)/2] {
			low = (low+high)/2 + 1
		} else {
			high = (low+high)/2 - 1
		}
	}
	return false
}

// 利用二分查找统计某个数字出现的次数
func GetCount(data []int, key int) (sum int) {
	// 定义左边的数量
	left := 0
	right := len(data) - 1
	location := 0
	for left <= right {
		mid := (left + right) / 2
		if data[mid] == key {
			// 记录找到的位置
			location = mid
			break
		} else if data[mid] > key {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	right = location
	left = location
	for data[right] == key && right >= 0 {
		right--
		sum++
	}
	for data[left] == key && left < len(data) {
		left++
		sum++
	}
	return sum - 1
}

func main() {
	data := []int{1, 3, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(GetCount(data, 3))
}
