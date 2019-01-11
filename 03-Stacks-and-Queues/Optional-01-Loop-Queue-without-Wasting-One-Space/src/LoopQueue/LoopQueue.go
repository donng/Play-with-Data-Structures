package main

import (
	"bytes"
	"fmt"
)

// 在这一版LoopQueue的实现中，我们将不浪费那1个空间：）
type LoopQueue struct {
	data  []interface{}
	front int
	tail  int
	size  int
}

func Constructor(capacity int) *LoopQueue {
	// 由于不浪费空间，所以data静态数组的大小是capacity
	// 而不是capacity + 1
	return &LoopQueue{
		data: make([]interface{}, capacity),
	}
}

func (this *LoopQueue) GetCapacity() int {
	return len(this.data)
}

func (this *LoopQueue) GetSize() int {
	return this.size
}

func (this *LoopQueue) IsEmpty() bool {
	// 注意，我们不再使用front和tail之间的关系来判断队列是否为空，而直接使用size
	return this.size == 0
}

// 入队
func (this *LoopQueue) Enqueue(e interface{}) {
	// 注意，我们不再使用front和tail之间的关系来判断队列是否为满，而直接使用size
	capacity := this.GetCapacity()
	if this.size == capacity {
		this.resize(capacity * 2)
	}
	this.data[this.tail] = e
	this.tail = (this.tail + 1) % len(this.data)
	this.size++
}

// 获得队列头部元素
func (this *LoopQueue) Dequeue() (e interface{}) {
	if this.IsEmpty() {
		panic("Cannot dequeue from empty queue")
	}

	e = this.data[this.front]
	this.data[this.front] = nil
	// 循环队列需要执行求余运算
	this.front = (this.front + 1) % len(this.data)
	this.size--
	if this.size == this.GetCapacity()/4 && this.size != 0 {
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
	length := len(this.data)
	newData := make([]interface{}, capacity)
	for i := 0; i < this.size; i++ {
		newData[i] = this.data[(i+this.front)%length]
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

func main() {
	queue := Constructor(10)
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)

		if i%3 == 2 {
			queue.Dequeue()
			fmt.Println(queue)
		}
	}
}
