package main

import "fmt"

// 搜索旋转排序的数组
func search(nums []int, target int) int {
	// 双指针
	i, j := 0, len(nums)-1
	for i <= j {
		if target == nums[i] {
			return i
		} else if target > nums[i] {
			i++
		} else if target == nums[j] {
			return j
		} else if target < nums[j] {
			j--
		}
	}
	return -1
}

// 二分法,更快
func search_two(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	mid := (left + right + 1) >> 1
	for left <= right {
		if nums[mid] == target {
			return mid
		}
		//左边升序
		if nums[left] < nums[mid] {
			// 如果目标范围
			if target <= nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { //右边是升序
			if target >= nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		mid = (left + right + 1) >> 1
	}

	return -1
}

func main() {
	data := []int{3, 1}
	fmt.Println(search_two(data, 1))
}
