package LinkedListQueue

import (
	"bytes"
	"fmt"
)

type node struct {
	e    interface{}
	next *node
}

func (n *node) String() string {
	return fmt.Sprint(n.e)
}

type LinkedListQueue struct {
	head *node
	tail *node
	size int
}

func Constructor() *LinkedListQueue {
	return &LinkedListQueue{}
}

func (this *LinkedListQueue) GetSize() int {
	return this.size
}

func (this *LinkedListQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *LinkedListQueue) Enqueue(e interface{}) {
	if this.tail == nil {
		this.tail = &node{e: e}
		this.head = this.tail
	} else {
		this.tail.next = &node{e: e}
		this.tail = this.tail.next
	}
	this.size++
}

func (this *LinkedListQueue) Dequeue() interface{} {
	if this.IsEmpty() {
		panic("Cannot dequeue from an empty queue.")
	}
	head := this.head
	this.head = this.head.next
	head.next = nil
	if head == nil {
		this.tail = nil
	}
	this.size--

	return head.e
}

func (this *LinkedListQueue) GetFront() interface{} {
	if this.IsEmpty() {
		panic("Cannot dequeue from an empty queue.")
	}

	return this.head.e
}

func (this *LinkedListQueue) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Queue: front ")

	cur := this.head
	for cur != nil {
		buffer.WriteString(fmt.Sprintf("%v ->", cur.e))
		cur = cur.next
	}
	buffer.WriteString("NULL tail")
	return buffer.String()
}

//func main()  {
//	queue := &LinkedListQueue{}
//	for i := 0; i < 10; i++ {
//		queue.Enqueue(i)
//		fmt.Println(queue)
//
//		if i % 3 == 2 {
//			queue.Dequeue()
//			fmt.Println(queue)
//		}
//	}
//}
