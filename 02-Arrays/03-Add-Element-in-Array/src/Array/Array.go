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

// 在第 index 个位置插入一个新元素 e
func (this *Array) Add(index int, e int) {
	if this.size == len(this.data) {
		panic("Add failed. Array is full.")
	}

	if index < 0 || index > this.GetCapacity() {
		panic("Add failed. Require index >= 0 and index <= size.")
	}

	for i := this.size - 1; i >= index; i-- {
		this.data[i+1] = this.data[i]
	}

	this.data[index] = e
	this.size++
}

// 向所有元素后添加一个新元素
func (this *Array) AddLast(e int) {
	//if this.size == len(this.data) {
	//	panic("AddLast failed,Array is full.")
	//}
	//
	//this.data[this.size] = e
	//this.size++
	this.Add(this.size, e)
}

// 向所有元素前添加一个新元素
func (this *Array) AddFirst(e int) {
	this.Add(0, e)
}
