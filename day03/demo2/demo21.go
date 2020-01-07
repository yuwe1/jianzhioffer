package main

import "fmt"

func Print1ToMaxOfDigits(n int) {
	if n <= 0 {
		return
	}
	number := make([]int, n)
	fmt.Println(number)
	for i := 0; i < 10; i++ {
		number[0] = i
		print1ToMaxOfDigitsRecursively(number, n, 0)
	}

}

func print1ToMaxOfDigitsRecursively(number []int, length int, index int) {
	if index == length-1 {
		printNumber(number)
		return
	}

	for i := 0; i < 10; i++ {
		
		number[index+1] = i
		print1ToMaxOfDigitsRecursively(number, length, index+1)
	}
}

func printNumber(number []int) {
	var isBeginning0 = true
	length := len(number)
	for i := 0; i < length; i++ {
		if isBeginning0 && number[i] != 0 {
			isBeginning0 = false
		}

		if !isBeginning0 {
			fmt.Printf("%d", number[i])
			if i == length-1 {
				fmt.Println()
			}
		}
	}
}

func main() {
	Print1ToMaxOfDigits(2)
}
