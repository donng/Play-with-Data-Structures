package Array

type Array struct {
	data []int // 自定义数组的数据
	size int   // 元素个数
}

// 获得指定容量的自定义数组
// 注：Go 中数组类型的长度是编译时指定，无法满足此需求，这里用 slice 代替
func GetArray(capacity int) *Array {
	a := &Array{}
	a.data = make([]int, capacity)
	a.size = 0
	return a
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
