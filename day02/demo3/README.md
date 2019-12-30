### 题目要求
求旋转数组的最小数字

### 什么是旋转数组
旋转数组是将一个有序数组按照某一个位置进行反转到其后，如：
```
有序数组：{123456}
按照4反转得{456123}
按照5反转得{561234}
按照1反转仍然为{123456}
```

### 解题思路
对于这个题，用遍历数组一个个比较去求出最小值即可，但是往往面试官不仅仅想看到这种解法，很明显，这题考察的是我们如何合理的利用旋转数组的特性。

通过上面的案例，我们可以了解到对于一个旋转数组，如果按照从第i（i>1）的位置反转，最小值的右边逐渐增大，左边为最大值开始从左往右的递减，我们可以总结出以下的规律：

假设最小的数为第i位，所以右边的数字递增，左边的数字递减。因此我们可以定义首尾指针，利用二分折半的方法进行寻找，思路如下:
```txt
1、定义i，j分别为首尾指针，mid为中间基准数
2、判断基准数和左指针的位置数值的大小，如果mid所指的数比较大，那么最小的数一定在mid的位置右边，如果比较小，那么最小的数一定在mid位置的左边（包含mid）
3、当左右指针相差为1时，即找到了最小的数
```

### 代码如下
```go
func selectMin(data []int) {
	// 定义首尾指针
	left, right := 0, len(data)-1
	if data[left] < data[right] {
		fmt.Println(data[left])
		return
	}
	for right >= left {
		if (right - left) == 1 {
			min := right
			fmt.Println(data[min])
			break
		}
		mid := (left + right) / 2
		// fmt.Println(data[mid])
		if data[mid] > data[left] {
			fmt.Println("左指针移动")
			// 说明mid位置为递增序列
			left = mid
		} else {
			fmt.Println("右指针移动")
			right = mid
		}
	}
}
```

### 结束了？
对于上面的思路其实我们少考虑了一种情况就是如果left mid right的位置所知的数相同，我们该如何判断，如10111，此时我们按照如上的思路将0左边的数字1会进行输出，再比如11101，我们按照如上的思路会将0右边的1进行输出。因此对于这种情况我们仍然需要按照顺序遍历的方法进行比较，完整代码如下：
```go
func selectMin(data []int) {
	// 定义首尾指针
	left, right := 0, len(data)-1
	if data[left] < data[right] {
		fmt.Println(data[left])
		return
	}
	for right >= left {
		if (right - left) == 1 {
			min := right
			fmt.Println(data[min])
			break
		}
		mid := (left + right) / 2
		// fmt.Println(data[mid])
		if data[mid] > data[left] {
			fmt.Println("左指针移动")
			// 说明mid位置为递增序列
			left = mid
		} else if data[mid] < data[left] {
			fmt.Println("右指针移动")
			right = mid
		} else {
			fmt.Println("特殊情况")
			// 特殊情况，当三个数字
			if data[left] == data[right] && data[right] == data[mid] {
				result := data[left]
				for i := left; i <= right; i++ {
					if data[i] < result {
						result = data[i]
					}
				}
				fmt.Println(result)
			}

			break
		}

	}
}
```
