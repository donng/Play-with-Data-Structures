package main

import (
	"bytes"
	"fmt"
)

type ArrayQueue struct {
	array *Array
}

func getArrayQueue(capacity int) (queue *ArrayQueue) {
	queue = &ArrayQueue{}
	queue.array = getArray(capacity)

	return
}

func (q *ArrayQueue) getSize() int {
	return q.array.getSize()
}

func (q *ArrayQueue) isEmpty() bool {
	return q.array.isEmpty()
}

func (q *ArrayQueue) getCapacity() int {
	return q.array.getCapacity()
}

func (q *ArrayQueue) enqueue(e interface{}) {
	q.array.addLast(e)
}

func (q *ArrayQueue) dequeue() interface{} {
	return q.array.removeFirst()
}

func (q *ArrayQueue) getFront() interface{} {
	return q.array.getFirst()
}

func (q *ArrayQueue) String() string {
	var buffer bytes.Buffer

	buffer.WriteString("Queue: ")
	buffer.WriteString("front [");
	for i := 0; i < q.array.getSize(); i++ {
		// fmt.Sprint 将 interface{} 类型转换为字符串
		buffer.WriteString(fmt.Sprint(q.array.data[i]))
		if i != (q.array.getSize() - 1) {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")

	return buffer.String()
}

func main()  {
	queue := getArrayQueue(20)
	for i := 0; i < 10; i++ {
		queue.enqueue(i)
		fmt.Println(queue)

		if i % 3 == 2 {
			queue.dequeue()
			fmt.Println(queue)
		}
	}
}
