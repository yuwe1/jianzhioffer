package main

import "fmt"

func getMax(data []int, target int) int {
	count := 0
	j := 0
	for i := 0; i < len(data); i++ {
		for i > 2 && (data[i]-data[j]) > target {
			j++
		}
		fmt.Println(i, ",", j)
		count += getCount(i - j)
		fmt.Println("count:", count)
	}
	return count % 99997867
}
func getCount(n int) int {
	return (n - 1) * n / 2
}
func main() {
	n := 0
	target := 0
	fmt.Scanf("%d %d\n", &n, &target)
	data := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &data[i])
	}
	fmt.Println(getMax(data, target))
}
