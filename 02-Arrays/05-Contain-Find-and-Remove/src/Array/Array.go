package Array

import (
	"bytes"
	"fmt"
	"strconv"
)

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

// 向所有元素后添加一个新元素
func (this *Array) AddLast(e int) {
	this.Add(this.size, e)
}

// 向所有元素前添加一个新元素
func (this *Array) AddFirst(e int) {
	this.Add(0, e)
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

// 获取 index 索引位置的元素
func (this *Array) Get(index int) int {
	if index < 0 || index >= this.size {
		panic("Get failed. Index is illegal.")
	}
	return this.data[index]
}

// 修改 index 索引位置的元素
func (this *Array) Set(index int, e int) {
	if index < 0 || index >= this.size {
		panic("Set failed. Index is illegal.")
	}
	this.data[index] = e
}

// 查找数组中是否有元素 e
func (this *Array) Contains(e int) bool {
	for i := 0; i < this.size; i++ {
		if this.data[i] == e {
			return true
		}
	}
	return false
}

// 查找数组中元素 e 所在的索引，不存在则返回 -1
func (this *Array) Find(e int) int {
	for i := 0; i < this.size; i++ {
		if this.data[i] == e {
			return i
		}
	}
	return -1
}

// 查找数组中元素 e 所有的索引组成的切片，不存在则返回 -1
func (this *Array) FindAll(e int) (indexes []int) {
	for i := 0; i < this.size; i++ {
		if this.data[i] == e {
			indexes = append(indexes, i)
		}
	}
	return
}

// 从数组中删除 index 位置的元素，返回删除的元素
func (this *Array) Remove(index int) (e int) {
	if index < 0 || index >= this.size {
		panic("Set failed. Index is illegal.")
	}

	e = this.data[index]
	for i := index + 1; i < this.size; i++ {
		this.data[i-1] = this.data[i]
	}
	this.size--
	return
}

// 从数组中删除第一个元素，返回删除的元素
func (this *Array) RemoveFirst() int {
	return this.Remove(0)
}

// 从数组中删除最后一个元素，返回删除的元素
func (this *Array) RemoveLast() int {
	return this.Remove(this.size - 1)
}

// 从数组中删除一个元素 e
func (this *Array) RemoveElement(e int) bool {
	index := this.Find(e)
	if index == -1 {
		return false
	}

	this.Remove(index)
	return true
}

// 从数组中删除所有元素 e
func (this *Array) RemoveAllElement(e int) bool {
	if this.Find(e) == -1 {
		return false
	}

	for i := 0; i < this.size; i++ {
		if this.data[i] == e {
			this.Remove(i)
		}
	}
	return true
}

// 重写 Array 的 string 方法
func (this *Array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", this.size, len(this.data)))
	buffer.WriteString("[")
	for i := 0; i < this.size; i++ {
		buffer.WriteString(strconv.Itoa(this.data[i]))
		if i != (this.size - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
