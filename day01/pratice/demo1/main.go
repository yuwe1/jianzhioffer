package main

import "fmt"

func insert_sort(data1, data2 []int) (data []int) {
	for i := 0; i < len(data1); i++ {
		data = append(data, data1[i])
	}
	for i := 0; i < len(data2); i++ {
		data = append(data, data2[i])
	}
	fmt.Println(data)
	// 定义两个指针指向尾部
	i, j := len(data1)-1, len(data)-1
	for i > 0 && j > 0 {
		if i == j {
			j = len(data) - 1
			i--
		}
		if data[i] > data[j] {
			data[i], data[j] = data[j], data[i]
			fmt.Println("交换", data)
		}
		j--
	}
	fmt.Println(data)
	return data
}

func merge_two(data1 []int, data2 []int) {
	i, j, k := 0, 0, 0
	var result []int
	result = make([]int, len(data1)+len(data2))
	for i < len(data1) && j < len(data2) {
		if data1[i] < data2[j] {
			result[k] = data1[i]
			k++
			i++
		} else {
			result[k] = data2[j]
			j++
			k++
		}

	}
	if i != len(data1) {
		for i < len(data1) {
			result[k] = data1[i]
			i++
			k++
		}
	}
	if j != len(data2) {
		for j < len(data2) {
			result[k] = data2[j]
			k++
			j++
		}
	}
	fmt.Println(result)
}
func main() {
	data1 := []int{1, 3, 5, 7, 8, 9, 10}
	data2 := []int{3, 4, 5, 6, 7, 8}
	// insert_sort(data1, data2)
	merge_two(data1, data2)
}
