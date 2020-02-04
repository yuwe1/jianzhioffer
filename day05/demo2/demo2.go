package main

// 定义一个堆
type heap struct {
	val      []int
	capacity int
	count    int
}

func NewHeap(capacity int) *heap {
	return &heap{
		val:      make([]int, capacity),
		capacity: capacity,
		count:    0,
	}
}

// 添加数字到堆
func (h *heap) addNode(val int) bool {
	if h.capacity == h.count {
		return false
	}
	h.val[h.count] = val
	i := h.count
	parent := (i - 1) / 2
	for parent >= 0 && (i-1) >= 0 {
		if h.val[parent] < h.val[i] {
			h.val[parent], h.val[i] = h.val[i], h.val[parent]
		}
		i = parent
		parent = (i - 1) / 2
	}
	h.count++
	return true
}

// 删除根节点
func (h *heap) MoveRoot() {
	h.val[h.count-1], h.val[0] = h.val[0], h.val[h.count-1]
	h.val[h.count-1] = 0
	h.count--
	// 调整堆的元素
	h.Adjust()

}

// 调整堆的元素
func (h *heap) Adjust() {
	top := 0
	//左边孩子的结点
	j := 2*top + 1
	for j < h.count {
		if j+1 < h.count && h.val[j+1] > h.val[j] {
			j++
		}
		if h.val[j] > h.val[top] {
			h.val[top], h.val[j] = h.val[j], h.val[top]
			top = j
			j = 2*top + 1
		} else {
			break
		}
	}

}

// 最小的k个数
// 使用最大堆存储最小的k个数
func GetKnum(nums []int, k int) {
	h := NewHeap(k)
	for i := 0; i < len(nums); i++ {
		if !h.addNode(nums[i]) && h.val[0] > nums[i] {
			// 删除根节点
			h.MoveRoot()
			h.addNode(nums[i])
		}

	}
}
func main() {
	nums := []int{4, 5, 1, 6, 2, 7, 3}

	GetKnum(nums, 4)
}
