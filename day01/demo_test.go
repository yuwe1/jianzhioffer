package day01

import (
	"fmt"
	demo1 "github.com/yuwe1/jianzhioffer/day01/demo1"
	"testing"
)

func TestFindNum(t *testing.T) {
	// 定义一个完整数组
	// var data [][4]int
	// data = make([][4]int, 4)
	// fmt.Println(data)[[0 0 0 0] [0 0 0 0] [0 0 0 0] [0 0 0 0]]
	// 初始化二维数组
	fmt.Println("测试用例1******")
	a := demo1.S2{
		demo1.S1{1, 2, 8, 9},
		demo1.S1{2, 4, 9, 12},
		demo1.S1{4, 7, 10, 13},
		demo1.S1{6, 8, 11, 15},
	}
	num1, num2 := 8, 3
	if demo1.FindNum(a, 4, 4, num1) == true && demo1.FindNum(a, 4, 4, num2) == false {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}
	fmt.Println("测试用例2********")
	a = demo1.S2{}
	if demo1.FindNum(a, 4, 4, num1) == false {
		fmt.Println("success")
	} else {
		fmt.Println("fail")
	}
}
