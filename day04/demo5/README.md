### 判断序列2是否是序列1形成栈的一种弹出序列
如序列1：1，2，3，4，5
序列2：4，5，3，2，1
对于以上弹出序列可用以下的压入弹出形成，因此是子序列
push(1),push(2),push(3),push(4),pop(4),push(5),pop(5),pop(3),pop(2),pop(1)

### 思路
对于以上的弹出思路，其实每一次的压入都要和弹出的序列进行比较，如果相等，则弹出，否则就压入剩余序列中的元素，如果可以通过弹出元素使的栈为空，则说明可以成为序列1的栈弹出序列
### 代码思路
- 准备两个指针指向两个序列
- 循环，循环的终止条件就是当弹出序列无法弹出的时候
- 当压入元素时，需要满足不超过边界，否则终止循环


```go
func IsPopOrder(pushdata []int, popdata []int) bool {
	j := 0
	i := 0
	for j < len(popdata) {
		if i < len(pushdata) {
			s.push(pushdata[i])
		} else {
			break
		}
		for j < len(popdata) && s.top.value == popdata[j] {
			fmt.Println(s.pop())
			j++
		}
		i++
	}
	if s.top == nil {
		return true
	}
	return false
}
```