package array

type Array struct {
	data []int
	size int
}

// New 构造函数，传入数组的容量capacity构造Array
func New(capacity int) Array {
	return Array{
		data: make([]int, capacity),
		size: 0,
	}
}

// GetCapacity 获取数组的容量
func (a *Array) GetCapacity() int {
	return len(a.data)
}

// GetSize 获得数组中的元素个数
func (a *Array) GetSize() int {
	return a.size
}

// IsEmpty 返回数组是否为空
func (a *Array) IsEmpty() bool {
	return a.size == 0
}
