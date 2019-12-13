package Array

type Array struct {
	data []int
	size int
}

// 构造函数，传入数组的容量capacity构造Array
func Constructor(capacity int) *Array {
	// size 默认为 0
	return &Array{
		data: make([]int, capacity),
	}
}

// 获取数组的容量
func (a *Array) GetCapacity() int {
	return len(a.data)
}

// 获得数组中的元素个数
func (a *Array) GetSize() int {
	return a.size
}

// 返回数组是否为空
func (a *Array) IsEmpty() bool {
	return a.size == 0
}
