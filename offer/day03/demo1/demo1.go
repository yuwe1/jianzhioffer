package main

//最大子序和，动态规划求解
func maxNum(data []int) {
	dp := make([]int, len(data))
	dp[0] = data[0]
	l := len(data)
	result := data[0]
	for i := 1; i < l; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		result := max(dp[i], result)
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
	data := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxNum(data)
}
