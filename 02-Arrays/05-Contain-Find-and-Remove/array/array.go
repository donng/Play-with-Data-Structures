package array

import (
	"bytes"
	"fmt"
	"strconv"
)

type Array struct {
	data []int
	size int
}

// New 构造函数，传入数组的容量capacity构造Array
func New(capacity int) *Array {
	return &Array{
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

// AddLast 向所有元素后添加一个新元素
func (a *Array) AddLast(e int) {
	a.Add(a.size, e)
}

// AddFirst 向所有元素前添加一个新元素
func (a *Array) AddFirst(e int) {
	a.Add(0, e)
}

// Add 在第 index 个位置插入一个新元素 e
func (a *Array) Add(index int, e int) {
	if a.size == len(a.data) {
		panic("add failed, array is full")
	}

	if index < 0 || index > a.size {
		panic("add failed, index out of range")
	}

	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}

	a.data[index] = e
	a.size++
}

// Get 获取 index 索引位置的元素
func (a *Array) Get(index int) int {
	if index < 0 || index >= a.size {
		panic("get failed, index out of range")
	}
	return a.data[index]
}

// Set 修改 index 索引位置的元素
func (a *Array) Set(index int, e int) {
	if index < 0 || index >= a.size {
		panic("set failed, index out of range")
	}
	a.data[index] = e
}

// Contains 查找数组中是否有元素 e
func (a *Array) Contains(e int) bool {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			return true
		}
	}
	return false
}

// Find 查找数组中元素 e 所在的索引，不存在则返回 -1
func (a *Array) Find(e int) int {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			return i
		}
	}
	return -1
}

// FindAll 查找数组中元素 e 所有的索引组成的切片，不存在则返回 -1
func (a *Array) FindAll(e int) (indexes []int) {
	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			indexes = append(indexes, i)
		}
	}
	return
}

// Remove 从数组中删除 index 位置的元素，返回删除的元素
func (a *Array) Remove(index int) int {
	if index < 0 || index >= a.size {
		panic("remove failed, index out of range")
	}

	e := a.data[index]
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}
	a.size--
	return e
}

// RemoveFirst 从数组中删除第一个元素，返回删除的元素
func (a *Array) RemoveFirst() int {
	return a.Remove(0)
}

// RemoveLast 从数组中删除最后一个元素，返回删除的元素
func (a *Array) RemoveLast() int {
	return a.Remove(a.size - 1)
}

// RemoveElement 从数组中删除一个元素 e
func (a *Array) RemoveElement(e int) bool {
	index := a.Find(e)
	if index == -1 {
		return false
	}

	a.Remove(index)
	return true
}

// RemoveAllElement 从数组中删除所有元素 e
func (a *Array) RemoveAllElement(e int) bool {
	if a.Find(e) == -1 {
		return false
	}

	for i := 0; i < a.size; i++ {
		if a.data[i] == e {
			a.Remove(i)
		}
	}
	return true
}

// 重写 Array 的 string 方法
func (a *Array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("array: size = %d, capacity = %d\n", a.size, len(a.data)))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		buffer.WriteString(strconv.Itoa(a.data[i]))
		if i != (a.size - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
