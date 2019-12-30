package main

import (
	"fmt"
)

func selectMin(data []int) {
	// 定义首尾指针
	left, right := 0, len(data)-1
	if data[left] < data[right] {
		fmt.Println(data[left])
		return
	}
	for right >= left {
		if (right - left) == 1 {
			min := right
			fmt.Println(data[min])
			break
		}
		mid := (left + right) / 2
		// fmt.Println(data[mid])
		if data[mid] > data[left] {
			fmt.Println("左指针移动")
			// 说明mid位置为递增序列
			left = mid
		} else if data[mid] < data[left] {
			fmt.Println("右指针移动")
			right = mid
		} else {
			fmt.Println("特殊情况")
			// 特殊情况，当三个数字
			if data[left] == data[right] && data[right] == data[mid] {
				result := data[left]
				for i := left; i <= right; i++ {
					if data[i] < result {
						result = data[i]
					}
				}
				fmt.Println(result)
			}

			break
		}

	}
}

func main() {
	// data := []int{3, 4, 5, 1, 2}
	// selectMin(data)
	// data2 := []int{1, 2, 3, 4, 5}
	// selectMin(data2)
	// data3 := []int{1, 0, 1, 1, 1}
	// selectMin(data3)
	data4 := []int{1, 1, 1, 0, 1}
	selectMin(data4)
}
