package Array

type Array struct {
	// 声明类型为 slice
	data []int
	size int
}

// 传入数组的容量 capacity 返回 Slice
// 注：在 Go 中不同长度的数组属于不同类型，所以这里使用 Slice
func GetArray(capacity int) (a *Array) {
	a = &Array{}
	a.data = make([]int, capacity)
	a.size = 0
	return
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


// 在第 index 个位置插入一个新元素 element
func (a *Array) Add(index int, element int) {
	if a.size == len(a.data) {
		panic("Add failed,Array is full")
	}

	if index < 0 || index > a.GetCapacity() {
		panic("Add failed,require index >= 0 and index <= a.cap")
	}

	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}

	a.data[index] = element
	a.size++
}


// 向所有元素后添加一个新元素
func (a *Array) AddLast(element int) {
	//if a.size == len(a.data) {
	//	panic("AddLast failed,Array is full.")
	//}
	//
	//a.data[a.size] = element
	//a.size++
	a.Add(a.size, element)
}

// 向所有元素前添加一个新元素
func (a *Array) AddFirst(element int) {
	a.Add(0, element)
}