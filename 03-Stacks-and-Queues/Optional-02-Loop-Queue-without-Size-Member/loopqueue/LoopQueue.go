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

func (lq *LoopQueue) GetCapacity() int {
	return len(lq.data) - 1
}

func (lq *LoopQueue) GetSize() int {
	// 注意此时getSize的逻辑:
	// 如果tail >= front，非常简单，队列中的元素个数就是tail - front
	// 如果tail < front，说明我们的循环队列"循环"起来了，此时，队列中的元素个数为：
	// tail - front + data.length
	// 画画图，看能不能理解为什么？
	//
	// 也可以理解成，此时，data中没有元素的数目为front - tail,
	// 整体元素个数就是 data.length - (front - tail) = data.length + tail - front
	if lq.tail >= lq.front {
		return lq.tail - lq.front
	} else {
		return lq.tail - lq.front + len(lq.data)
	}
}

func (lq *LoopQueue) IsEmpty() bool {
	return lq.front == lq.tail
}

// 入队
func (lq *LoopQueue) Enqueue(e interface{}) {
	if (lq.tail+1)%len(lq.data) == lq.front {
		lq.resize(lq.GetCapacity() * 2)
	}
	lq.data[lq.tail] = e
	lq.tail = (lq.tail + 1) % len(lq.data)
}

// 获得队列头部元素
func (lq *LoopQueue) Dequeue() (e interface{}) {
	if lq.IsEmpty() {
		panic("cannot dequeue from empty queue")
	}

	e = lq.data[lq.front]
	lq.data[lq.front] = nil
	// 循环队列需要执行求余运算
	lq.front = (lq.front + 1) % len(lq.data)
	if lq.GetSize() == lq.GetCapacity()/4 && lq.GetSize() != 0 {
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
	newData := make([]interface{}, capacity+1)
	sz := lq.GetSize()
	for i := 0; i < sz; i++ {
		newData[i] = lq.data[(i+lq.front)%len(lq.data)]
	}
	lq.data = newData
	lq.front = 0
	lq.tail = sz
}

func (lq *LoopQueue) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Queue: size = %d, capacity = %d\n", lq.GetSize(), lq.GetCapacity()))
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
