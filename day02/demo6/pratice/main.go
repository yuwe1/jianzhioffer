package main

import "fmt"

// 判断一个整数是否是2的整数次方2^0 2^1 2^3 2^4
func IntergerMod2(data int) bool {
	if (data-1)&data == 0 {
		return true
	}
	return false
}

// 将一个整数转换成其它整数，需要将其二进制转换几位
func GetChangeCount(data, changedata int) int {
	data = data ^ changedata
	count := 0
	for data != 0 {
		count++
		data = data & (data - 1)
	}
	return count
}
func main() {
	fmt.Println(IntergerMod2(10))
	fmt.Println(IntergerMod2(9))
	fmt.Println(IntergerMod2(-9))
	fmt.Println(IntergerMod2(0))
	fmt.Println(GetChangeCount(10, 13))
}
