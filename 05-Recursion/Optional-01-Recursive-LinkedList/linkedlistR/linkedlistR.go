package linkedlistR

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

func New() *LinkedListR {
	return &LinkedListR{}
}

// 获取链表中的元素个数
func (lr *LinkedListR) GetSize() int {
	return lr.size
}

// 返回链表是否为空
func (lr *LinkedListR) IsEmpty() bool {
	return lr.size == 0
}

// 在链表的index(0-based)位置添加新的元素e
func (lr *LinkedListR) Add(index int, e interface{}) {
	if index < 0 || index > lr.size {
		panic("Add failed. Illegal index.")
	}

	lr.head = lr.add(lr.head, index, e)
	lr.size++
}

// 在以node为头结点的链表的index位置插入元素e，递归算法
func (lr *LinkedListR) add(n *Node, index int, e interface{}) *Node {
	if index == 0 {
		return &Node{e, n}
	}

	n.Next = lr.add(n.Next, index-1, e)
	return n
}

// 在链表头添加新的元素e
func (lr *LinkedListR) AddFirst(e interface{}) {
	lr.Add(0, e)
}

// 在链表末尾添加新的元素e
func (lr *LinkedListR) AddLast(e interface{}) {
	lr.Add(lr.size, e)
}

// 获得链表的第index(0-based)个位置的元素
func (lr *LinkedListR) Get(index int) interface{} {
	if index < 0 || index > (lr.size-1) {
		panic("Get failed. Illegal index.")
	}

	return lr.get(lr.head, index)
}

// 在以node为头结点的链表中，找到第index个元素，递归算法
func (lr *LinkedListR) get(n *Node, index int) interface{} {
	if index == 0 {
		return n.e
	}

	return lr.get(n.Next, index-1)
}

// 获得链表的第一个元素
func (lr *LinkedListR) GetFirst() interface{} {
	return lr.Get(0)
}

// 获得链表的最后一个元素
func (lr *LinkedListR) GetLast() interface{} {
	return lr.Get(lr.size - 1)
}

// 修改链表的第index(0-based)个位置的元素为e
func (lr *LinkedListR) Set(index int, e interface{}) {
	if index < 0 || index > (lr.size-1) {
		panic("Set failed. Illegal index.")
	}

	lr.set(lr.head, index, e)
}

// 修改以node为头结点的链表中，第index(0-based)个位置的元素为e，递归算法
func (lr *LinkedListR) set(node *Node, index int, e interface{}) {
	if index == 0 {
		node.e = e
		return
	}

	lr.set(node.Next, index-1, e)
}

// 查找链表中是否有元素e
func (lr *LinkedListR) Contains(e interface{}) bool {
	return lr.contains(lr.head, e)
}

// 在以node为头结点的链表中，查找是否存在元素e，递归算法
func (lr *LinkedListR) contains(n *Node, e interface{}) bool {
	if n == nil {
		return false
	}
	if n.e == e {
		return true
	}

	return lr.contains(n.Next, e)
}

// 从链表中删除index(0-based)位置的元素, 返回删除的元素
func (lr *LinkedListR) Remove(index int) interface{} {
	head, e := lr.remove(lr.head, index)
	lr.size--
	// 只剩头节点时，head 为 nil
	lr.head = head
	return e
}

// 从以node为头结点的链表中，删除第index位置的元素，递归算法
// 返回值包含两个元素，删除后的链表头结点和删除的值：）
func (lr *LinkedListR) remove(n *Node, index int) (*Node, interface{}) {
	if index == 0 {
		// 最终返回待删除节点的下一个节点和待删除节点的值
		return n.Next, n.e
	}

	nextNode, e := lr.remove(n.Next, index-1)
	// 最内层递归：待删除节点的上一个节点，设置 Next 为待删除节点的下一个节点
	n.Next = nextNode
	return n, e
}

// 从链表中删除第一个元素, 返回删除的元素
func (lr *LinkedListR) RemoveFirst() interface{} {
	return lr.Remove(0)
}

// 从链表中删除最后一个元素, 返回删除的元素
func (lr *LinkedListR) RemoveLast() interface{} {
	return lr.Remove(lr.size - 1)
}

// 从链表中删除元素e
func (lr *LinkedListR) RemoveElement(e interface{}) {
	lr.removeElement(lr.head, e)
}

// 从以node为头结点的链表中，删除元素e，递归算法
func (lr *LinkedListR) removeElement(n *Node, e interface{}) *Node {
	if n == nil {
		return nil
	}
	if n.e == e {
		lr.size--
		return n.Next
	}
	n.Next = lr.removeElement(n.Next, e)
	return n
}

func (lr *LinkedListR) String() string {
	buffer := bytes.Buffer{}

	cur := lr.head
	for cur != nil {
		buffer.WriteString(fmt.Sprintf("%v ->", cur.e))
		cur = cur.Next
	}

	buffer.WriteString("NULL")
	return buffer.String()
}
