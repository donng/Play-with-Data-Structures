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

func (lq *LoopQueue) GetCapacity() int {
	return len(lq.data)
}

func (lq *LoopQueue) GetSize() int {
	return lq.size
}

func (lq *LoopQueue) IsEmpty() bool {
	// 注意，我们不再使用front和tail之间的关系来判断队列是否为空，而直接使用size
	return lq.size == 0
}

// 入队
func (lq *LoopQueue) Enqueue(e interface{}) {
	// 注意，我们不再使用front和tail之间的关系来判断队列是否为满，而直接使用size
	capacity := lq.GetCapacity()
	if lq.size == capacity {
		lq.resize(capacity * 2)
	}
	lq.data[lq.tail] = e
	lq.tail = (lq.tail + 1) % len(lq.data)
	lq.size++
}

// 获得队列头部元素
func (lq *LoopQueue) Dequeue() (e interface{}) {
	if lq.IsEmpty() {
		panic("Cannot dequeue from empty queue")
	}

	e = lq.data[lq.front]
	lq.data[lq.front] = nil
	// 循环队列需要执行求余运算
	lq.front = (lq.front + 1) % len(lq.data)
	lq.size--
	if lq.size == lq.GetCapacity()/4 && lq.size != 0 {
		lq.resize(lq.GetCapacity() / 2)
	}

	return
}

// 查看队列头部元素
func (lq *LoopQueue) GetFront() interface{} {
	if lq.IsEmpty() {
		panic("Queue is empty")
	}

	return lq.data[lq.front]
}

func (lq *LoopQueue) resize(capacity int) {
	length := len(lq.data)
	newData := make([]interface{}, capacity)
	for i := 0; i < lq.size; i++ {
		newData[i] = lq.data[(i+lq.front)%length]
	}
	lq.data = newData
	lq.front = 0
	lq.tail = lq.size
}

func (lq *LoopQueue) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Queue: size = %d, capacity = %d\n", lq.size, lq.GetCapacity()))
	buffer.WriteString("front [")
	for i := lq.front; i != lq.tail; i = (i + 1) % len(lq.data) {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprint(lq.data[i]))
		if (i+1)%len(lq.data) != lq.tail {
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
