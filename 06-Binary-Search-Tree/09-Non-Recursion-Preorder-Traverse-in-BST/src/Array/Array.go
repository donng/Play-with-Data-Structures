package Array

import (
	"bytes"
	"fmt"
)

type Array struct {
	// 声明类型为 slice
	data []interface{}
	size int
}

// 传入数组的容量 capacity 返回 Slice
// 注：在 Go 中不同长度的数组属于不同类型，所以这里使用 Slice
func GetArray(capacity int) (a *Array) {
	a = &Array{}
	a.data = make([]interface{}, capacity)
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
func (a *Array) Add(index int, element interface{}) {
	if index < 0 || index > a.GetCapacity() {
		panic("Add failed,require index >= 0 and index <= a.cap")
	}

	if a.size == len(a.data) {
		a.resize(2 * a.size)
	}

	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}

	a.data[index] = element
	a.size++
}

// 向所有元素后添加一个新元素
func (a *Array) AddLast(element interface{}) {
	//if a.size == len(a.data) {
	//	panic("AddLast failed,Array is full.")
	//}
	//
	//a.data[a.size] = element
	//a.size++
	a.Add(a.size, element)
}

// 向所有元素前添加一个新元素
func (a *Array) AddFirst(element interface{}) {
	a.Add(0, element)
}

// 获取 index 索引位置的元素
func (a *Array) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("Get failed,Index is illegal.")
	}
	return a.data[index]
}

func (a *Array) GetLast() interface{} {
	return a.Get(a.size - 1)
}

func (a *Array) GetFirst() interface{} {
	return a.Get(0)
}

// 修改 index 索引位置的元素
func (a *Array) Set(index int, element interface{}) {
	if index < 0 || index >= a.size {
		panic("Set failed,Index is illegal.")
	}
	a.data[index] = element
}

// 查找数组中是否有元素 element
func (a *Array) Contains(element interface{}) bool {
	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			return true
		}
	}

	return false
}

// 查找数组中元素 element 所在的索引，不存在则返回 -1
func (a *Array) Find(element interface{}) int {
	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			return i
		}
	}

	return -1
}

// 查找数组中元素 element 所有的索引组成的切片，不存在则返回 -1
func (a *Array) FindAll(element interface{}) (indexes []int) {
	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			indexes = append(indexes, i)
		}
	}

	return
}

// 从数组中删除 index 位置的元素，返回删除的元素
func (a *Array) Remove(index int) (element interface{}) {
	if index < 0 || index >= a.size {
		panic("Set failed,Index is illegal.")
	}

	element = a.data[index]
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}
	a.size--
	a.data[a.size] = nil //loitering object != memory leak

	if a.size == len(a.data)/2 {
		a.resize(len(a.data) / 2)
	}
	return
}

// 从数组中删除第一个元素，返回删除的元素
func (a *Array) RemoveFirst() interface{} {
	return a.Remove(0)
}

// 从数组中删除最后一个元素，返回删除的元素
func (a *Array) RemoveLast() interface{} {
	return a.Remove(a.size - 1)
}

// 从数组中删除一个元素 element
func (a *Array) RemoveElement(element interface{}) bool {
	index := a.Find(element)
	if index == -1 {
		return false
	}

	a.Remove(index)
	return true
}

// 从数组中删除所有元素 element
func (a *Array) RemoveAllElement(element interface{}) bool {
	if a.Find(element) == -1 {
		return false
	}

	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			a.Remove(i)
		}
	}
	return true
}

// 为数组扩容
func (a *Array) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity)
	for i := 0; i < a.size; i++ {
		newData[i] = a.data[i]
	}

	a.data = newData
}

// 重写 Array 的 string 方法
func (a *Array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", a.size, len(a.data)))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprint(a.data[i]))
		if i != (a.size - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
