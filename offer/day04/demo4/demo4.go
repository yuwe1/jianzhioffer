package main

import "fmt"

// 数组的交集不排序
func GetSame(data1, data2 []int) (nums []int) {
	if len(data1) == 0 || len(data2) == 0 {
		return
	}
	var m = map[int]int{}
	fmt.Println(len(m))
	for i := 0; i < len(data1); i++ {
		m[data1[i]]++
	}
	fmt.Println(len(m))
	for j := 0; j < len(data2); j++ {
		if m[data2[j]] > 0 {
			m[data2[j]]--
			nums = append(nums, data2[j])
		}
	}
	return
}

func main() {
	data1 := []int{2, 2}
	data2 := []int{1, 4, 2, 1, 3, 2}
	fmt.Println(GetSame(data1, data2))
}
