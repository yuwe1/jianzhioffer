package main

import "fmt"

// 最长上升子序列

func longnums(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	result := 1
	for i := 0; i < l; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
		result = max(result, dp[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	data := []int{4, 10, 4, 3, 8, 9}
	fmt.Println(longnums(data))
}
