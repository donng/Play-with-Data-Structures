package main

import (
	"bytes"
	"fmt"
)

// 在这一版本的实现中，我们完全不使用size，只使用front和tail来完成LoopQueue的所有逻辑：）
type LoopQueue struct {
	data  []interface{}
	front int
	tail  int
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
	// 注意此时getSize的逻辑:
	// 如果tail >= front，非常简单，队列中的元素个数就是tail - front
	// 如果tail < front，说明我们的循环队列"循环"起来了，此时，队列中的元素个数为：
	// tail - front + data.length
	// 画画图，看能不能理解为什么？
	//
	// 也可以理解成，此时，data中没有元素的数目为front - tail,
	// 整体元素个数就是 data.length - (front - tail) = data.length + tail - front
	if this.tail >= this.front {
		return this.tail - this.front
	} else {
		return this.tail - this.front + len(this.data)
	}
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
	if this.GetSize() == this.GetCapacity()/4 && this.GetSize() != 0 {
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
	sz := this.GetSize()
	for i := 0; i < sz; i++ {
		newData[i] = this.data[(i+this.front)%len(this.data)]
	}
	this.data = newData
	this.front = 0
	this.tail = sz
}

func (this *LoopQueue) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Queue: size = %d, capacity = %d\n", this.GetSize(), this.GetCapacity()))
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
