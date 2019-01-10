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
func (this *Array) GetCapacity() int {
	return len(this.data)
}

// 获得数组中的元素个数
func (this *Array) GetSize() int {
	return this.size
}

// 返回数组是否为空
func (this *Array) IsEmpty() bool {
	return this.size == 0
}
