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
