package main

import "fmt"

// 思路1:n>0
func Power_one(data float64, n int) float64 {
	sum := 1.0
	for i := 1; i <= n; i++ {
		sum *= data
	}
	return sum
}

//思路2：0的0次方没有意义，返回1或者0都可以
func Power_two(data float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	sign := 0
	if n < 0 {
		n = -n
		sign = 1
	}
	sum := Power_one(data, n)
	if sign == 1 && sum != 0 {
		return 1.0 / sum
	}
	return sum
}

// 在浮点数中判断两个数是否相等时，不能直接用==，因为在计算机中表示小数都是有误差的
func equal(num1, num2 float64) bool {
	if num1-num2 > -0.0000001 && num1-num2 < 0.0000001 {
		return true
	}
	return false
}

//思路3：代码优化
func Power_three(data float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	sign := 0
	if n < 0 {
		n = -n
		sign = 1
	}
	sum := Power_one(data, n)
	if sign == 1 && !equal(sum, 0.0) {
		return 1.0 / sum
	}
	return sum
}

// 更快的方法求一个正整数次方
func Power_one__two(data float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return data
	}

	sum := Power_one__two(data, n>>1)
	sum *= sum
	//判断n为奇数还是偶数
	if n&1 == 1 {
		//	奇数
		sum *= data
	}
	return sum

}

// 降低时间复杂度优化
func Power_four(data float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	sign := 0
	if n < 0 {
		n = -n
		sign = 1
	}
	sum := Power_one__two(data, n)
	if sign == 1 && !equal(sum, 0.0) {
		return 1.0 / sum
	}
	return sum
}
func main() {
	// fmt.Println(Power_two(2, 3))
	// fmt.Println(Power_two(-1, 3))
	// fmt.Println(Power_two(-2, -3))
	// fmt.Println(Power_two(0, -3))
	fmt.Println(Power_four(2, 0))
	fmt.Println(Power_four(0, 0))
	fmt.Println(Power_four(2, -2))
	fmt.Println(Power_four(2, 2))
	fmt.Println(Power_four(2, -1))
}
