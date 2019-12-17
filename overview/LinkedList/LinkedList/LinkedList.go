package LinkedList

import (
	"bytes"
	"fmt"
)

// 单链表

type Node struct {
	e    interface{}
	next *Node
}

type LinkedList struct {
	dummyHead *Node
	size      int
}

func Constructor() *LinkedList {
	return &LinkedList{
		dummyHead: &Node{},
	}
}

func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedList) GetSize() int {
	return l.size
}

func (l *LinkedList) Add(index int, e interface{}) {
	if index < 0 || index > l.size {
		panic("Add failed. Index must >= 0 and <= size.")
	}

	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	prev.next = &Node{e: e, next: prev.next}
	l.size++
}

func (l *LinkedList) AddFirst(e interface{}) {
	l.Add(0, e)
}

func (l *LinkedList) AddLast(e interface{}) {
	l.Add(l.size, e)
}

func (l *LinkedList) Remove(index int) interface{} {
	if index < 0 || index >= l.size {
		panic("Remove failed. Index must >= 0 and < size.")
	}
	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	// cur 为待删除节点
	cur := prev.next
	prev.next = cur.next
	cur.next = nil
	l.size--

	return cur.e
}

func (l *LinkedList) RemoveFirst() interface{} {
	return l.Remove(0)
}

func (l *LinkedList) RemoveLast() interface{} {
	return l.Remove(l.size - 1)
}

func (l *LinkedList) Get(index int) interface{} {
	if index < 0 || index >= l.size {
		panic("Get failed. Index must >= 0 and < size.")
	}
	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.e
}

func (l *LinkedList) GetFirst() interface{} {
	return l.Get(0)
}

func (l *LinkedList) GetLast() interface{} {
	return l.Get(l.size - 1)
}

func (l *LinkedList) Contains(e interface{}) bool {
	cur := l.dummyHead.next
	for cur != nil {
		if cur.e == e {
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *LinkedList) Set(index int, e interface{}) {
	if index < 0 || index >= l.size {
		panic("Get failed. Index must >= 0 and < size.")
	}

	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.e = e
}

func (l *LinkedList) String() string {
	var buffer bytes.Buffer

	cur := l.dummyHead.next
	for cur.next != nil {
		buffer.WriteString(fmt.Sprintf("%s->", cur.e))
		cur = cur.next
	}
	buffer.WriteString(fmt.Sprintf("%s->nil, size: %d", cur.e, l.size))
	return buffer.String()
}

func main() {
	linkList := Constructor()
	linkList.AddFirst("are")
	fmt.Println(linkList)

	linkList.AddFirst("where")
	linkList.AddLast("you")
	linkList.AddLast("form")
	linkList.RemoveLast()
	fmt.Println(linkList)

	linkList.AddLast("from")

	fmt.Println(linkList)
}
