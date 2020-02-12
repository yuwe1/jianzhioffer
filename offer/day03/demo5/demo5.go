package main

import "fmt"

// 两个有序数组的合并
func Merage(nums1 []int, nums2 []int) {
	m := len(nums1)
	n := len(nums2)
	p1, p2 := m-1, n-1
	p := m + n - 1
	for i := 0; i < n; i++ {
		nums1 = append(nums1, 0)
	}
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] < nums2[p2] {
			nums1[p] = nums2[p2]
			p2--
			p--
		} else {
			nums1[p] = nums1[p1]
			p--
			p1--
		}
	}
	if p2 > 0 {
		for i := p2; i >= 0; i-- {
			nums1[p] = nums2[i]
			p--
		}
	}
	fmt.Println(nums1)
}

func main() {
	data1 := []int{1, 3, 9, 10}
	data2 := []int{2, 4, 6, 8}
	Merage(data1, data2)
}
