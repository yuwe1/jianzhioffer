package main

import "fmt"

func movingCount(threshold, rows, cols int) int {
	sum := 0
	for i := 0; i < rows; i++ {
		maxj := cols - 1
		for maxj >= 0 && i+maxj > threshold {
			maxj--
		}
		fmt.Println(sum)
		sum += maxj + 1
	}
	return sum
}
func main() {
	fmt.Println(movingCount(9, 20, 25))
}
