package LoopQueue

import (
	"bytes"
	"fmt"
)

type LoopQueue struct {
	data  []interface{}
	front int
	tail  int
	size  int
}

func Constructor(capacity int) *LoopQueue {
	return &LoopQueue{
		data: make([]interface{}, capacity+1),
	}
}

func (this *LoopQueue) GetCapacity() int {
	return len(this.data) - 1
}

func (this *LoopQueue) GetSize() int {
	return this.size
}

func (this *LoopQueue) IsEmpty() bool {
	return this.front == this.tail
}

// 入队
func (this *LoopQueue) Enqueue(e interface{}) {
	if (this.tail+1)%len(this.data) == this.front {
		this.resize(this.GetCapacity() * 2)
	}
	this.data[this.tail] = e
	this.tail = (this.tail + 1) % len(this.data)
	this.size++
}

// 获得队列头部元素
func (this *LoopQueue) Dequeue() (e interface{}) {
	if this.IsEmpty() {
		panic("cannot dequeue from empty queue")
	}

	e = this.data[this.front]
	this.data[this.front] = nil
	// 循环队列需要执行求余运算
	this.front = (this.front + 1) % len(this.data)
	this.size--
	if this.size == this.GetCapacity()/4 && this.GetCapacity()/2 != 0 {
		this.resize(this.GetCapacity() / 2)
	}

	return
}

// 查看队列头部元素
func (this *LoopQueue) GetFront() interface{} {
	if this.IsEmpty() {
		panic("Queue is empty")
	}

	return this.data[this.front]
}

func (this *LoopQueue) resize(capacity int) {
	newData := make([]interface{}, capacity+1)
	for i := 0; i < this.size; i++ {
		newData[i] = this.data[(i+this.front)%len(this.data)]
	}
	this.data = newData
	this.front = 0
	this.tail = this.size
}

func (this *LoopQueue) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Queue: size = %d, capacity = %d\n", this.size, this.GetCapacity()))
	buffer.WriteString("front [")
	for i := this.front; i != this.tail; i = (i + 1) % len(this.data) {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprint(this.data[i]))
		if (i+1)%len(this.data) != this.tail {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")

	return buffer.String()
}
