package ArrayQueue

import (
	"Play-with-Data-Structures/03-Stacks-and-Queues/08-Queues-Comparison/src/Array"
	"bytes"
	"fmt"
)

// 数组队列的局限性：出队列的操作时间复杂度为 n
type ArrayQueue struct {
	array *Array.Array
}

func Constructor(capacity int) *ArrayQueue {
	return &ArrayQueue{
		array: Array.Constructor(capacity),
	}
}

func (this *ArrayQueue) GetSize() int {
	return this.array.GetSize()
}

func (this *ArrayQueue) IsEmpty() bool {
	return this.array.IsEmpty()
}

func (this *ArrayQueue) GetCapacity() int {
	return this.array.GetCapacity()
}

func (this *ArrayQueue) Enqueue(e interface{}) {
	this.array.AddLast(e)
}

func (this *ArrayQueue) Dequeue() interface{} {
	return this.array.RemoveFirst()
}

func (this *ArrayQueue) GetFront() interface{} {
	return this.array.GetFirst()
}

func (this *ArrayQueue) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Queue: ")
	buffer.WriteString("front [")
	for i := 0; i < this.array.GetSize(); i++ {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprint(this.array.Get(i)))
		if i != (this.array.GetSize() - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")

	return buffer.String()
}
