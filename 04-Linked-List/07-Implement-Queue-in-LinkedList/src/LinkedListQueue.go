package main

import (
	"fmt"
	"bytes"
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

func (l *LinkedListQueue) GetSize() int {
	return l.size
}

func (l *LinkedListQueue) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedListQueue) Enqueue(e interface{}) {
	if l.tail == nil {
		l.tail = &node{e: e}
		l.head = l.tail
	} else {
		l.tail.next = &node{e: e}
		l.tail = l.tail.next
	}
	l.size++
}

func (l *LinkedListQueue) Dequeue() interface{} {
	if l.IsEmpty() {
		panic("Cannot dequeue from an empty queue.")
	}
	head := l.head
	l.head = l.head.next
	head.next = nil
	if head == nil {
		l.tail = nil
	}
	l.size--

	return head.e
}

func (l *LinkedListQueue) GetFront() interface{} {
	if l.IsEmpty() {
		panic("Cannot dequeue from an empty queue.")
	}

	return l.head.e
}

func (l *LinkedListQueue) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Queue: front ")

	cur := l.head
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
