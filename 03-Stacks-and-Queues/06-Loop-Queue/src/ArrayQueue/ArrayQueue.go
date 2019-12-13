package ArrayQueue

import (
	"Play-with-Data-Structures/03-Stacks-and-Queues/06-Loop-Queue/src/Array"
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

func (aq *ArrayQueue) GetSize() int {
	return aq.array.GetSize()
}

func (aq *ArrayQueue) IsEmpty() bool {
	return aq.array.IsEmpty()
}

func (aq *ArrayQueue) GetCapacity() int {
	return aq.array.GetCapacity()
}

func (aq *ArrayQueue) Enqueue(e interface{}) {
	aq.array.AddLast(e)
}

func (aq *ArrayQueue) Dequeue() interface{} {
	return aq.array.RemoveFirst()
}

func (aq *ArrayQueue) GetFront() interface{} {
	return aq.array.GetFirst()
}

func (aq *ArrayQueue) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Queue: ")
	buffer.WriteString("front [")
	for i := 0; i < aq.array.GetSize(); i++ {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprint(aq.array.Get(i)))
		if i != (aq.array.GetSize() - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")

	return buffer.String()
}
