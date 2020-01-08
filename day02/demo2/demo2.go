package main

import "fmt"

// {4, 2, 1, 10, 9, 10, 2, 0}
func bubbleSort(data []int) {
	// 外循环控制次数
	for i := 0; i < len(data)-1; i++ {
		//内循环进行比较
		for j := 0; j < len(data)-1-i; j++ {
			// fmt.Println(data[j], "与", data[j+1], "作比较")
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	fmt.Println(data)
}

// 选择排序
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

// 直接插入排序扩展数组写法
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

// 直接插入排序简单写法
func Insert_Sort(data []int) {
	// 外循环控制待插入的数字
	for i := 1; i < len(data); i++ {
		for j := i; j > 0; j-- {
			if data[j] > data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
}

// 折半插入排序写法——简单
func Birary_Insert_Sort(data []int) {
	//外循环控制待插入的数字
	for i := 1; i < len(data); i++ {
		// 内循环寻找位置
		// 待插入的记录
		temp := data[i]
		low := 0
		high := i - 1
		for low <= high {
			mid := (low + high) / 2
			if data[mid] > temp {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		for j := i - 1; j >= high+1; j-- {
			data[j+1] = data[j]
		}
		data[high+1] = temp
	}
}

// 折半插入排序扩展数组写法
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
// 归并排序
func Merge_Sort(arg []int, res []int)[]int {
	// 比较这两个临时数组进行合并
	i := 0
	// 即指向第二个数组的起始位置
	j := 0
	var result []int
	for i != len(arg) && j != len(res) {
		if arg[i] > res[j] {
			result = append(result, res[j])
			j++
		} else {
			result = append(result, arg[i])
			i++
		}
	}
	if j == len(res) {
		for i != len(arg) {
			result = append(result, arg[i])
			i++
		}
	} else {
		for j != len(res) {
			result = append(result, res[j])
			j++
		}
	}
	fmt.Println(result)
	return result

}
func Merge(args []int){
	// 定义一个二维数组
	var result [][]int
	for i :=0;i<len(args);i++{
		temp := args[i:i+1]
		result = append(result,temp)
	}
	
	for len(result) >= 2{
		var te [][]int
		for i:=0;i<len(result);i+=2{
			temp :=Merge_Sort(result[i],result[i+1])
			te = append(te,temp)
		}
		result = te
	}	
	fmt.Println(result)
}
// 快速排序
func Quick_Sort(data []int,left,right int){
	// 定义基准数的位置
	p := left
	i,j :=left,right
	for i<=j{
		if data[p] <=data[j] && p <=j{
			j--
		}
		if p <= j{
			data[p],data[j] = data[j],data[p]
			p = j
		}
		if data[p]>=data[i] && i <=p{
			i++
		}
		if p >=i{
			data[p],data[i] = data[i],data[p]
			p = i
		}
	}
	if (p-left)>1{
		Quick_Sort(data,left,p-1)
	}
	if (right-p)>1{
		Quick_Sort(data,p+1,right)
	}
}
// 对公司所有员工年龄排序
func SortAges(ages []int){
	// 设置最大年龄
	oldage := 10
	// 用来存储年龄的次数
	var tempage []int
	for i :=0;i<=oldage;i++{
		tempage = append(tempage,0)
	}
	for i:=0;i<len(ages);i++{
		// 统计每个年龄出现的次数
		
		tempage[ages[i]]++
	}
	fmt.Println(tempage)
	index :=0
	for i:=0;i<=oldage;i++{
		for j :=0;j < tempage[i];j++{
			ages[index] = i
			index++
		}
	}
}
func main() {
	data := []int{4, 2, 1, 10, 9, 10, 2, 0}
	Quick_Sort(data,0,7)
	fmt.Println(data)
	// bubbleSort(data)
	// Select_Sort(data)
	// for i := 0; i < len(data); i++ {
	// 	direct_insert(data[:i], data[i])
	// }
	// fmt.Println(data)
	// for i := 0; i < len(data); i++ {
	// 	Birary_Serect_Sort(data[:i], data[i])
	// }
	// Birary_Serect_Sort(data[:0], data[0])
	// Birary_Serect_Sort(data[:1], data[1])
	// Birary_Serect_Sort(data[:2], data[2])
	// Birary_Serect_Sort(data[:3], data[3])
	// Birary_Serect_Sort(data[:4], data[4])
	// Birary_Serect_Sort(data[:5], data[5])
	// Birary_Serect_Sort(data[:6], data[6])
	// Birary_Serect_Sort(data[:7], data[7])
	// Insert_Sort(data)
	// Birary_Insert_Sort(data)

	// fmt.Println(data)
	// args :=[]int{4}
	// res :=[]int{1}
	// Merge_Sort(args,res)
	// Quick_Sort(data,0,len(data)-1)
	// Merge(data)
	// SortAges(data)
	// fmt.Println(data)

}
