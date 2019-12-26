package main

import "fmt"

func MaoPao(data []int) {
	// 外循环控制次数
	for i := 0; i < len(data)-1; i++ {
		//内循环进行比较
		for j := 0; j < len(data)-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	fmt.Println(data)
}
func Select_Sort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
	fmt.Println(data)
}

// 直接插入排序
func direct_insert(data []int, key int) {
	// 获取已经排好序列的数组的长度
	length := len(data)
	data = append(data, key)
	// 定义尾指针
	tail := length - 1
	// 定义result的末尾
	right := length

	for tail >= 0 && data[tail] > key {

		data[tail], data[right] = data[right], data[tail]
		tail--
		right--

	}
}

func Birary_Serect_Sort(data []int, key int) {
	if len(data) == 0 {
		data = append(data, key)
		fmt.Println(data)
		return
	}
	left := 0
	right := len(data) - 1
	// 扩充数组
	data = append(data, key)
	tail := len(data) - 1
	location := 0
	for left <= right {
		if right == left && data[right] <= key {
			location = right + 1
			break
		} else if right == left && data[left] >= key {
			location = left
			break
		} else {
			if (right - left) == 1 {
				if data[left] <= key && data[right] >= key {
					location = left + 1
					break
				}
				break
			} else {
				mid := (left + right) / 2
				if data[mid] < key {
					left = mid + 1
				} else if data[mid] > key {
					right = mid - 1
				}
			}
		}
	}
	// fmt.Println(data)
	right = tail - 1
	for right >= location && right >= 0 {
		data[tail], data[right] = data[right], data[tail]
		right--
		tail--
	}
	fmt.Println(data)
}
func main() {
	data := []int{4, 2, 1, 10, 9, 10, 2, 0}
	// MaoPao(data)
	// Select_Sort(data)
	// for i := 0; i < len(data); i++ {
	// 	direct_insert(data[:i], data[i])
	// }
	// fmt.Println(data)
	// for i := 0; i < len(data); i++ {
	// 	Birary_Serect_Sort(data[:i], data[i])
	// }
	Birary_Serect_Sort(data[:0], data[0])
	Birary_Serect_Sort(data[:1], data[1])
	Birary_Serect_Sort(data[:2], data[2])
	Birary_Serect_Sort(data[:3], data[3])
	Birary_Serect_Sort(data[:4], data[4])
	Birary_Serect_Sort(data[:5], data[5])
	Birary_Serect_Sort(data[:6], data[6])
	Birary_Serect_Sort(data[:7], data[7])
	fmt.Println(data)
}
