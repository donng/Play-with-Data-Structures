package LinkedListR

import (
	"bytes"
	"fmt"
)

type Node struct {
	E    interface{}
	Next *Node
}

type LinkedListR struct {
	head *Node
	size int
}

type E interface{}

func GetLinkedListR() *LinkedListR {
	return &LinkedListR{}
}

// 获取链表中的元素个数
func (l *LinkedListR) GetSize() int {
	return l.size
}

// 返回链表是否为空
func (l *LinkedListR) IsEmpty() bool {
	return l.size == 0
}

// 在链表的index(0-based)位置添加新的元素e
func (l *LinkedListR) Add(index int, e E) {
	if index < 0 || index > l.size {
		panic("Add failed. Illegal index.")
	}

	l.head = l.add(l.head, index, e)
	l.size++
}

// 在以node为头结点的链表的index位置插入元素e，递归算法
func (l *LinkedListR) add(node *Node, index int, e E) *Node {
	if index == 0 {
		return &Node{e, node}
	}

	node.Next = l.add(node.Next, index-1, e)
	return node
}

// 在链表头添加新的元素e
func (l *LinkedListR) AddFirst(e E) {
	l.Add(0, e)
}

// 在链表末尾添加新的元素e
func (l *LinkedListR) AddLast(e E) {
	l.Add(l.size, e)
}

// 获得链表的第index(0-based)个位置的元素
func (l *LinkedListR) Get(index int) E {
	if index < 0 || index > (l.size-1) {
		panic("Get failed. Illegal index.")
	}

	return l.get(l.head, index)
}

// 在以node为头结点的链表中，找到第index个元素，递归算法
func (l *LinkedListR) get(node *Node, index int) E {
	if index == 0 {
		return node.E
	}

	return l.get(node.Next, index-1)
}

// 获得链表的第一个元素
func (l *LinkedListR) GetFirst() E {
	return l.Get(0)
}

// 获得链表的最后一个元素
func (l *LinkedListR) GetLast() E {
	return l.Get(l.size - 1)
}

// 修改链表的第index(0-based)个位置的元素为e
func (l *LinkedListR) Set(index int, e E) {
	if index < 0 || index > (l.size-1) {
		panic("Set failed. Illegal index.")
	}

	l.set(l.head, index, e)
}

// 修改以node为头结点的链表中，第index(0-based)个位置的元素为e，递归算法
func (l *LinkedListR) set(node *Node, index int, e E) {
	if index == 0 {
		node.E = e
		return
	}

	l.set(node.Next, index-1, e)
}

// 查找链表中是否有元素e
func (l *LinkedListR) Contains(e E) bool {
	return l.contains(l.head, e)
}

// 在以node为头结点的链表中，查找是否存在元素e，递归算法
func (l *LinkedListR) contains(node *Node, e E) bool {
	if node == nil {
		return false
	}
	if node.E == e {
		return true
	}

	return l.contains(node.Next, e)
}

// 从链表中删除index(0-based)位置的元素, 返回删除的元素
func (l *LinkedListR) Remove(index int) E {
	head, e := l.remove(l.head, index)
	l.size--
	// 只剩头节点时，head 为 nil
	l.head = head
	return e
}

// 从以node为头结点的链表中，删除第index位置的元素，递归算法
// 返回值包含两个元素，删除后的链表头结点和删除的值：）
func (l *LinkedListR) remove(node *Node, index int) (*Node, E) {
	if index == 0 {
		// 最终返回待删除节点的下一个节点和待删除节点的值
		return node.Next, node.E
	}

	nextNode, e := l.remove(node.Next, index-1)
	// 最内层递归：待删除节点的上一个节点，设置 Next 为待删除节点的下一个节点
	node.Next = nextNode
	return node, e
}

// 从链表中删除第一个元素, 返回删除的元素
func (l *LinkedListR) RemoveFirst() E {
	return l.Remove(0)
}

// 从链表中删除最后一个元素, 返回删除的元素
func (l *LinkedListR) RemoveLast() E {
	return l.Remove(l.size - 1)
}

// 从链表中删除元素e
func (l *LinkedListR) RemoveElement(e E) {
	l.removeElement(l.head, e)
}

// 从以node为头结点的链表中，删除元素e，递归算法
func (l *LinkedListR) removeElement(node *Node, e E) *Node {
	if node == nil {
		return nil
	}
	if node.E == e {
		l.size--
		return node.Next
	}
	node.Next = l.removeElement(node.Next, e)
	return node
}

func (l *LinkedListR) String() string {
	buffer := bytes.Buffer{}

	cur := l.head
	for cur != nil {
		buffer.WriteString(fmt.Sprintf("%v ->", cur.E))
		cur = cur.Next
	}

	buffer.WriteString("NULL")
	return buffer.String()
}
