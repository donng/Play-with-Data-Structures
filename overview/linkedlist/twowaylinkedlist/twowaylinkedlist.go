package main

import (
	"bytes"
	"fmt"
)

// 单链表
// 添加和删除时，遇到尾节点的情况需要特殊处理

type Node struct {
	e          interface{}
	prev, next *Node
}

type TwoWayLinkedList struct {
	head *Node
	size int
}

func New() *TwoWayLinkedList {
	return &TwoWayLinkedList{}
}

func (l *TwoWayLinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *TwoWayLinkedList) GetSize() int {
	return l.size
}

func (l *TwoWayLinkedList) Add(index int, e interface{}) {
	if index < 0 || index > l.size {
		panic("add failed, index is out of range")
	}

	if index == 0 {
		if l.head == nil {
			l.head = &Node{e: e}
		} else {
			retNode := &Node{
				e:    e,
				next: l.head,
			}
			l.head.prev = retNode
			l.head = retNode
		}

	} else {
		prev := l.head
		for i := 0; i < index-1; i++ {
			prev = prev.next
		}

		newNode := &Node{
			e:    e,
			prev: prev,
			next: prev.next,
		}
		// 判断是否是尾节点
		if prev.next != nil {
			prev.next.prev = newNode
		}
		prev.next = newNode
	}
	l.size++
}

func (l *TwoWayLinkedList) AddFirst(e interface{}) {
	l.Add(0, e)
}

func (l *TwoWayLinkedList) AddLast(e interface{}) {
	l.Add(l.size, e)
}

func (l *TwoWayLinkedList) Remove(index int) interface{} {
	if index < 0 || index >= l.size {
		panic("remove failed, index is out of range")
	}

	var ele *Node
	if index == 0 {
		// 删除头结点
		l.head.next.prev = nil
		ele = l.head
		l.head = l.head.next
	} else {
		prev := l.head
		for i := 0; i < index-1; i++ {
			prev = prev.next
		}
		ele = prev.next
		// 修改待删除节点的后一个节点 prev 指向其前一个节点
		if prev.next.next != nil {
			prev.next.next.prev = prev
		}
		prev.next = prev.next.next
	}
	ele.prev = nil
	ele.next = nil
	l.size--

	return ele.e
}

func (l *TwoWayLinkedList) RemoveFirst() interface{} {
	return l.Remove(0)
}

func (l *TwoWayLinkedList) RemoveLast() interface{} {
	return l.Remove(l.size - 1)
}

func (l *TwoWayLinkedList) Get(index int) interface{} {
	if index < 0 || index >= l.size {
		panic("get failed, index is out of range")
	}
	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.e
}

func (l *TwoWayLinkedList) GetFirst() interface{} {
	return l.Get(0)
}

func (l *TwoWayLinkedList) GetLast() interface{} {
	return l.Get(l.size - 1)
}

func (l *TwoWayLinkedList) Contains(e interface{}) bool {
	cur := l.head
	for cur != nil {
		if cur.e == e {
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *TwoWayLinkedList) Set(index int, e interface{}) {
	if index < 0 || index >= l.size {
		panic("set failed, index is out of range")
	}

	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.e = e
}

func (l *TwoWayLinkedList) String() string {
	var buffer bytes.Buffer

	cur := l.head
	for cur.next != nil {
		buffer.WriteString(fmt.Sprintf("%s->", cur.e))
		cur = cur.next
	}
	buffer.WriteString(fmt.Sprintf("%s->nil, size: %d", cur.e, l.size))
	return buffer.String()
}

func main() {
	linkList := New()
	linkList.AddFirst("are")
	fmt.Println(linkList)

	linkList.AddFirst("where")
	fmt.Println(linkList)
	linkList.AddLast("you")
	fmt.Println(linkList)
	linkList.AddLast("form")
	linkList.RemoveLast()
	fmt.Println(linkList)

	linkList.AddLast("from")

	fmt.Println(linkList)
}
