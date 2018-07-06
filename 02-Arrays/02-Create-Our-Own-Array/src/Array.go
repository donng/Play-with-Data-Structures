package main

type Array struct {
	// 声明类型为 slice
	data []int
	size int
}

// 传入数组的容量 capacity 返回 Slice
// 注：在 Go 中不同长度的数组属于不同类型，所以这里使用 Slice
func getArray(capacity int) (a *Array) {
	a.data = make([]int, capacity)
	a.size = 0
	return
}

// 获取数组的容量
func (a *Array) getCapacity() int  {
	return len(a.data)
}

// 获得数组中的元素个数
func (a *Array) getSize() int {
	return a.size
}

// 返回数组是否为空
func (a *Array) isEmpty() bool {
	return a.size == 0
}

