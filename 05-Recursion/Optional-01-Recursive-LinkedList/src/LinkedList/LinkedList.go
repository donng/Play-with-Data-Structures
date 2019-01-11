package LinkedList

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

func Constructor() *LinkedList {
	return &LinkedList{
		dummyHead: &Node{},
	}
}

// 获取链表中的元素个数
func (this *LinkedList) GetSize() int {
	return this.size
}

// 返回链表是否为空
func (this *LinkedList) IsEmpty() bool {
	return this.size == 0
}

// 在链表的index(0-based)位置添加新的元素e
// 在链表中不是一个常用的操作，练习用：）
func (this *LinkedList) Add(index int, e interface{}) {
	if index < 0 || index > this.size {
		panic("Add failed. Illegal index.")
	}

	// 获得待插入节点的前一个节点
	prev := this.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	prev.next = &Node{e, prev.next}
	this.size++
}

// 在链表头添加新的元素e
func (this *LinkedList) AddFirst(e interface{}) {
	this.Add(0, e)
}

// 在链表末尾添加新的元素e
func (this *LinkedList) AddLast(e interface{}) {
	this.Add(this.size, e)
}

// 获得链表的第index(0-based)个位置的元素
// 在链表中不是一个常用的操作，练习用：）
func (this *LinkedList) Get(index int) interface{} {
	if index < 0 || index >= this.size {
		panic("Add failed. Illegal index.")
	}

	cur := this.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.e
}

// 获得链表的第一个元素
func (this *LinkedList) GetFirst() interface{} {
	return this.Get(0)
}

// 获得链表的最后一个元素
func (this *LinkedList) GetLast() interface{} {
	return this.Get(this.size - 1)
}

// 修改链表的第index(0-based)个位置的元素为e
// 在链表中不是一个常用的操作，练习用：）
func (this *LinkedList) Set(index int, e interface{}) {
	if index < 0 || index >= this.size {
		panic("Set failed. Illegal index.")
	}

	cur := this.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.e = e
}

// 查找链表是否存在元素e
func (this *LinkedList) Contains(e interface{}) bool {
	cur := this.dummyHead.next

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
func (this *LinkedList) Remove(index int) interface{} {
	if index < 0 || index >= this.size {
		panic("Remove failed. Index is illegal.")
	}

	// prev 是待删除元素的前一个元素
	prev := this.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}

	retNode := prev.next
	prev.next = retNode.next
	retNode.next = nil
	this.size--

	return retNode.e
}

// 从链表中删除第一个元素，返回删除的元素
func (this *LinkedList) RemoveFirst() {
	this.Remove(0)
}

// 从链表中删除最后一个元素，返回删除的元素
func (this *LinkedList) RemoveLast() {
	this.Remove(this.size - 1)
}

// 从链表中删除元素e
func (this *LinkedList) RemoveElement(e interface{}) {
	prev := this.dummyHead
	for prev.next != nil {
		if prev.next.e == e {
			delNode := prev.next
			prev.next = delNode.next
			delNode.next = nil
			this.size--

			break
		}
		prev = prev.next
	}
}

func (this *LinkedList) String() string {
	buffer := bytes.Buffer{}
	cur := this.dummyHead.next
	for cur != nil {
		buffer.WriteString(fmt.Sprint(cur.e) + "->")
		cur = cur.next
	}

	buffer.WriteString("NULL")

	return buffer.String()
}
