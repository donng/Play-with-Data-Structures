package linkedlist

import (
	"bytes"
	"fmt"
)

type Node struct {
	e    interface{}
	next *Node
}

func (n *Node) String() string {
	return fmt.Sprint(n.e)
}

type LinkedList struct {
	dummyHead *Node // 虚拟头结点，不计入size
	size      int
}

func New() *LinkedList {
	return &LinkedList{
		dummyHead: &Node{},
	}
}

// 获取链表中的元素个数
func (l *LinkedList) GetSize() int {
	return l.size
}

// 返回链表是否为空
func (l *LinkedList) IsEmpty() bool {
	return l.size == 0
}

// 在链表的index(0-based)位置添加新的元素e
// 在链表中不是一个常用的操作，练习用：）
func (l *LinkedList) Add(index int, e interface{}) {
	if index < 0 || index > l.size {
		panic("add failed, index is out of range")
	}

	// 获得待插入节点的前一个节点
	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	prev.next = &Node{e, prev.next}
	l.size++
}

// 在链表头添加新的元素e
func (l *LinkedList) AddFirst(e interface{}) {
	l.Add(0, e)
}

// 在链表末尾添加新的元素e
func (l *LinkedList) AddLast(e interface{}) {
	l.Add(l.size, e)
}

// 获得链表的第index(0-based)个位置的元素
// 在链表中不是一个常用的操作，练习用：）
func (l *LinkedList) Get(index int) interface{} {
	if index < 0 || index >= l.size {
		panic("get failed, index is out of range")
	}

	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.e
}

// 获得链表的第一个元素
func (l *LinkedList) GetFirst() interface{} {
	return l.Get(0)
}

// 获得链表的最后一个元素
func (l *LinkedList) GetLast() interface{} {
	return l.Get(l.size - 1)
}

// 修改链表的第index(0-based)个位置的元素为e
// 在链表中不是一个常用的操作，练习用：）
func (l *LinkedList) Set(index int, e interface{}) {
	if index < 0 || index >= l.size {
		panic("set failed, index is out of range")
	}

	cur := l.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.e = e
}

// 查找链表是否存在元素e
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

// 从链表中删除index(0-based)位置的元素，返回删除的元素
// 在链表中不是一个常用的操作，练习用：）
func (l *LinkedList) Remove(index int) interface{} {
	if index < 0 || index >= l.size {
		panic("remove failed, index is out of range")
	}

	// prev 是待删除元素的前一个元素
	prev := l.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	retNode := prev.next
	prev.next = retNode.next
	retNode.next = nil
	l.size--

	return retNode.e
}

// 从链表中删除第一个元素，返回删除的元素
func (l *LinkedList) RemoveFirst() {
	l.Remove(0)
}

// 从链表中删除最后一个元素，返回删除的元素
func (l *LinkedList) RemoveLast() {
	l.Remove(l.size - 1)
}

// 从链表中删除元素e
func (l *LinkedList) RemoveElement(e interface{}) {
	prev := l.dummyHead
	for prev.next != nil {
		if prev.next.e == e {
			delNode := prev.next
			prev.next = delNode.next
			delNode.next = nil
			l.size--

			break
		}
		prev = prev.next
	}
}

func (l *LinkedList) String() string {
	buffer := bytes.Buffer{}
	cur := l.dummyHead.next
	for cur != nil {
		buffer.WriteString(fmt.Sprint(cur.e) + "->")
		cur = cur.next
	}

	buffer.WriteString("NULL")

	return buffer.String()
}
