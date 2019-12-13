package Array

type Array struct {
	data []int
	size int
}

// 构造函数，传入数组的容量capacity构造Array
func Constructor(capacity int) *Array {
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

// 在第 index 个位置插入一个新元素 e
func (a *Array) Add(index int, e int) {
	if a.size == len(a.data) {
		panic("Add failed. Array is full.")
	}

	if index < 0 || index > a.size {
		panic("Add failed. Require index >= 0 and index <= size.")
	}

	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}

	a.data[index] = e
	a.size++
}

// 向所有元素后添加一个新元素
func (a *Array) AddLast(e int) {
	//if a.size == len(a.data) {
	//	panic("AddLast failed,Array is full.")
	//}
	//
	//a.data[a.size] = e
	//a.size++
	a.Add(a.size, e)
}

// 向所有元素前添加一个新元素
func (a *Array) AddFirst(e int) {
	a.Add(0, e)
}
