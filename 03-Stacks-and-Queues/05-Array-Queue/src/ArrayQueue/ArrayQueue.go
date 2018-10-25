package main

import (
	"Play-with-Data-Structures/03-Stacks-and-Queues/05-Array-Queue/src/Array"
	"bytes"
	"fmt"
)

// 数组队列的局限性：出队列的操作时间复杂度为 n
type ArrayQueue struct {
	array *Array.Array
}

func getArrayQueue(capacity int) (queue *ArrayQueue) {
	queue = &ArrayQueue{}
	queue.array = Array.GetArray(capacity)

	return
}

func (q *ArrayQueue) GetSize() int {
	return q.array.GetSize()
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.array.IsEmpty()
}

func (q *ArrayQueue) GetCapacity() int {
	return q.array.GetCapacity()
}

func (q *ArrayQueue) Enqueue(e interface{}) {
	q.array.AddLast(e)
}

func (q *ArrayQueue) Dequeue() interface{} {
	return q.array.RemoveFirst()
}

func (q *ArrayQueue) GetFront() interface{} {
	return q.array.GetFirst()
}

func (q *ArrayQueue) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Queue: ")
	buffer.WriteString("front [");
	for i := 0; i < q.array.GetSize(); i++ {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprint(q.array.Get(i)))
		if i != (q.array.GetSize() - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")

	return buffer.String()
}

func main()  {
	queue := getArrayQueue(20)
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)

		if i % 3 == 2 {
			queue.Dequeue()
			fmt.Println(queue)
		}
	}
}
