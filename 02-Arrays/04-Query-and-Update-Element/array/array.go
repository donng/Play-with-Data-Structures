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
		panic("add failed, required index >= 0 and index <= size")
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
		panic("get failed, index is illegal")
	}
	return a.data[index]
}

// Set 修改 index 索引位置的元素
func (a *Array) Set(index int, e int) {
	if index < 0 || index >= a.size {
		panic("get failed, index is illegal")
	}
	a.data[index] = e
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
