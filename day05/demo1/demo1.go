package main

import (
	"fmt"
)

// 数组中出现次数超过一半的数字

// 利用快速排序
func MoreThanHalfNum(nums []int, left, right int) {
	i, j := left, right
	// 基准值
	mid := left
	for i <= j {
		for mid <= j && nums[j] >= nums[mid] {
			j--
		}
		if mid <= j {
			nums[mid], nums[j] = nums[j], nums[mid]
			mid = j
		}
		for mid >= i && nums[i] <= nums[mid] {
			i++
		}
		if mid >= i {
			nums[i], nums[mid] = nums[mid], nums[i]
			mid = i
		}
	}
	if (right - mid) > 1 {
		MoreThanHalfNum(nums, mid+1, right)
	}
	if (mid - left) > 1 {
		MoreThanHalfNum(nums, left, mid-1)
	}
}

// 利用次数(对数字的范围有限制)
func MoreThanHalfNum_sec(nums []int) {
	var count [100]int
	max := len(nums) / 2
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
		if count[nums[i]] > max {
			fmt.Println(nums[i])
			break
		}
	}

}
func main() {
	nums := []int{1, 2, 3, 2, 2, 2, 2, 2, 5, 4, 2}
	// MoreThanHalfNum(nums, 0, 10)
	// fmt.Println(nums[len(nums)/2])
	MoreThanHalfNum_sec(nums)
}
