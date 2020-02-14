package main

// 只出现一次的数字
func singleNumber(nums []int) int {
	num := 0
	for i := 0; i < len(nums); i++ {
		num = num ^ nums[i]
	}
	return num
}
func main() {

}
