package main

import (
	"bytes"
	"fmt"
)

// 循环链表
// 循环链表终止条件，n.next = l.head。循环链表为一个环形，对头结点的操作需要特殊处理，例如：删除头结点需要修改尾节点指向第二个节点位置。

type Node struct {
	e    interface{}
	next *Node
}

type LoopLinkedList struct {
	head *Node
	size int
}

func New() *LoopLinkedList {
	return &LoopLinkedList{}
}

func (l *LoopLinkedList) IsEmpty() bool {
	return l.size == 0
}

func (l *LoopLinkedList) GetSize() int {
	return l.size
}

func (l *LoopLinkedList) Add(index int, e interface{}) {
	if index < 0 || index > l.size {
		panic("add failed, index is out of range")
	}

	if l.head == nil {
		n := &Node{e: e}
		l.head = n
		n.next = n
	} else {
		cur := l.head
		// 已存在头结点，考虑
		if index == 0 {
			l.head = &Node{
				e:    e,
				next: l.head,
			}
			// 添加头结点，需要修改尾节点指针指向新的头结点
			for cur.next != l.head.next {
				cur = cur.next
			}
			cur.next = l.head
		} else {
			// 插入节点是获取index位置的上一个节点
			prev := l.head
			for i := 0; i < index-1; i++ {
				prev = prev.next
			}
			prev.next = &Node{
				e:    e,
				next: prev.next,
			}
		}
	}

	l.size++
}

func (l *LoopLinkedList) AddFirst(e interface{}) {
	l.Add(0, e)
}

func (l *LoopLinkedList) AddLast(e interface{}) {
	l.Add(l.size, e)
}

func (l *LoopLinkedList) Remove(index int) interface{} {
	if index < 0 || index >= l.size {
		panic("remove failed, index is out of range")
	}

	var ele *Node
	if index == 0 {
		ele = l.head
		// 修改尾节点指针指向新的头结点
		cur := l.head
		for cur.next != l.head {
			cur = cur.next
		}
		cur.next = l.head.next
		l.head = l.head.next
	} else {
		// 除头节点外，其他节点的处理方式一致（包括尾节点）
		prev := l.head
		for i := 0; i < index-1; i++ {
			prev = prev.next
		}
		ele = prev.next
		prev.next = ele.next
	}

	ele.next = nil
	l.size--

	return ele.e
}

func (l *LoopLinkedList) RemoveFirst() interface{} {
	return l.Remove(0)
}

func (l *LoopLinkedList) RemoveLast() interface{} {
	return l.Remove(l.size - 1)
}

func (l *LoopLinkedList) Get(index int) interface{} {
	if index < 0 || index >= l.size {
		panic("get failed, index is out of range")
	}
	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.e
}

func (l *LoopLinkedList) GetFirst() interface{} {
	return l.Get(0)
}

func (l *LoopLinkedList) GetLast() interface{} {
	return l.Get(l.size - 1)
}

func (l *LoopLinkedList) Contains(e interface{}) bool {
	cur := l.head
	for cur != nil {
		if cur.e == e {
			return true
		}
		cur = cur.next
	}
	return false
}

func (l *LoopLinkedList) Set(index int, e interface{}) {
	if index < 0 || index >= l.size {
		panic("get failed, index is out of range")
	}

	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.e = e
}

func (l *LoopLinkedList) String() string {
	var buffer bytes.Buffer

	cur := l.head
	for cur.next != l.head {
		buffer.WriteString(fmt.Sprintf("%s->", cur.e))
		cur = cur.next
	}
	buffer.WriteString(fmt.Sprintf("%s->%s(HEAD), size: %d", cur.e, cur.next.e, l.size))
	return buffer.String()
}

func main() {
	linkList := New()
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
