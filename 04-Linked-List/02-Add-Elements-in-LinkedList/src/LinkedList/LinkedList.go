package LinkedList

import "fmt"

type Node struct {
	e    interface{}
	next *Node
}

func (n *Node) String() string {
	return fmt.Sprint(n.e)
}

type LinkedList struct {
	head *Node
	size int
}

// 获取链表中的元素个数
func (this *LinkedList) GetSize() int {
	return this.size
}

// 返回链表是否为空
func (this *LinkedList) IsEmpty() bool {
	return this.size == 0
}

// 在链表头添加新的元素e
func (this *LinkedList) AddFirst(e interface{}) {
	//Node := &Node{e: e}
	//Node.next = this.head
	//this.head = Node
	this.head = &Node{e, this.head}
	this.size++
}

// 在链表的index(0-based)位置添加新的元素e
// 在链表中不是一个常用的操作，练习用：）
func (this *LinkedList) Add(index int, e interface{}) {
	if index < 0 || index > this.size {
		panic("Add failed. Illegal index.")
	}

	if index == 0 {
		this.AddFirst(e)
	} else {
		// 获得待插入节点的前一个节点
		prev := this.head
		for i := 0; i < index-1; i++ {
			prev = prev.next
		}

		// 插入新节点
		//Node := &Node{e: e, next: prev.next}
		//prev.next = Node
		prev.next = &Node{e, prev.next}
		this.size++
	}
}

// 在链表末尾添加新的元素e
func (this *LinkedList) AddLast(e interface{}) {
	this.Add(this.size, e)
}
