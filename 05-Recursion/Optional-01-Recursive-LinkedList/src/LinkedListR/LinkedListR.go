package LinkedListR

import (
	"bytes"
	"fmt"
)

type Node struct {
	e    interface{}
	Next *Node
}

type LinkedListR struct {
	head *Node
	size int
}

func Constructor() *LinkedListR {
	return &LinkedListR{}
}

// 获取链表中的元素个数
func (this *LinkedListR) GetSize() int {
	return this.size
}

// 返回链表是否为空
func (this *LinkedListR) IsEmpty() bool {
	return this.size == 0
}

// 在链表的index(0-based)位置添加新的元素e
func (this *LinkedListR) Add(index int, e interface{}) {
	if index < 0 || index > this.size {
		panic("Add failed. Illegal index.")
	}

	this.head = this.add(this.head, index, e)
	this.size++
}

// 在以node为头结点的链表的index位置插入元素e，递归算法
func (this *LinkedListR) add(n *Node, index int, e interface{}) *Node {
	if index == 0 {
		return &Node{e, n}
	}

	n.Next = this.add(n.Next, index-1, e)
	return n
}

// 在链表头添加新的元素e
func (this *LinkedListR) AddFirst(e interface{}) {
	this.Add(0, e)
}

// 在链表末尾添加新的元素e
func (this *LinkedListR) AddLast(e interface{}) {
	this.Add(this.size, e)
}

// 获得链表的第index(0-based)个位置的元素
func (this *LinkedListR) Get(index int) interface{} {
	if index < 0 || index > (this.size-1) {
		panic("Get failed. Illegal index.")
	}

	return this.get(this.head, index)
}

// 在以node为头结点的链表中，找到第index个元素，递归算法
func (this *LinkedListR) get(n *Node, index int) interface{} {
	if index == 0 {
		return n.e
	}

	return this.get(n.Next, index-1)
}

// 获得链表的第一个元素
func (this *LinkedListR) GetFirst() interface{} {
	return this.Get(0)
}

// 获得链表的最后一个元素
func (this *LinkedListR) GetLast() interface{} {
	return this.Get(this.size - 1)
}

// 修改链表的第index(0-based)个位置的元素为e
func (this *LinkedListR) Set(index int, e interface{}) {
	if index < 0 || index > (this.size-1) {
		panic("Set failed. Illegal index.")
	}

	this.set(this.head, index, e)
}

// 修改以node为头结点的链表中，第index(0-based)个位置的元素为e，递归算法
func (this *LinkedListR) set(node *Node, index int, e interface{}) {
	if index == 0 {
		node.e = e
		return
	}

	this.set(node.Next, index-1, e)
}

// 查找链表中是否有元素e
func (this *LinkedListR) Contains(e interface{}) bool {
	return this.contains(this.head, e)
}

// 在以node为头结点的链表中，查找是否存在元素e，递归算法
func (this *LinkedListR) contains(n *Node, e interface{}) bool {
	if n == nil {
		return false
	}
	if n.e == e {
		return true
	}

	return this.contains(n.Next, e)
}

// 从链表中删除index(0-based)位置的元素, 返回删除的元素
func (this *LinkedListR) Remove(index int) interface{} {
	head, e := this.remove(this.head, index)
	this.size--
	// 只剩头节点时，head 为 nil
	this.head = head
	return e
}

// 从以node为头结点的链表中，删除第index位置的元素，递归算法
// 返回值包含两个元素，删除后的链表头结点和删除的值：）
func (this *LinkedListR) remove(n *Node, index int) (*Node, interface{}) {
	if index == 0 {
		// 最终返回待删除节点的下一个节点和待删除节点的值
		return n.Next, n.e
	}

	nextNode, e := this.remove(n.Next, index-1)
	// 最内层递归：待删除节点的上一个节点，设置 Next 为待删除节点的下一个节点
	n.Next = nextNode
	return n, e
}

// 从链表中删除第一个元素, 返回删除的元素
func (this *LinkedListR) RemoveFirst() interface{} {
	return this.Remove(0)
}

// 从链表中删除最后一个元素, 返回删除的元素
func (this *LinkedListR) RemoveLast() interface{} {
	return this.Remove(this.size - 1)
}

// 从链表中删除元素e
func (this *LinkedListR) RemoveElement(e interface{}) {
	this.removeElement(this.head, e)
}

// 从以node为头结点的链表中，删除元素e，递归算法
func (this *LinkedListR) removeElement(n *Node, e interface{}) *Node {
	if n == nil {
		return nil
	}
	if n.e == e {
		this.size--
		return n.Next
	}
	n.Next = this.removeElement(n.Next, e)
	return n
}

func (this *LinkedListR) String() string {
	buffer := bytes.Buffer{}

	cur := this.head
	for cur != nil {
		buffer.WriteString(fmt.Sprintf("%v ->", cur.e))
		cur = cur.Next
	}

	buffer.WriteString("NULL")
	return buffer.String()
}
