package main

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

func IsPalindrome(head *Node) bool {
	if head == nil || head.next == nil {
		return true
	}

	var prev *Node
	slow := head
	fast := head

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		// 慢节点将要移动到的下一个节点
		next := slow.next
		slow.next = prev
		prev = slow
		slow = next
	}

	// 当 fast 不为 nil，说明 slow 在奇数个节点的中间位置，移动到下一个位置
	// a<-b->a  prev 在a，slow在b，需要将slow移动到第二个a
	if fast != nil {
		slow = slow.next
	}

	for slow != nil {
		if slow.e != prev.e {
			return false
		}
		slow = slow.next
		prev = prev.next
	}
	return true
}

func main() {
	testStr1 := []string{"a", "b", "c", "b", "a"} // true
	testStr2 := []string{"a", "b", "c", "a", "b"} // false

	linkList1 := Constructor()
	for _, s := range testStr1 {
		linkList1.AddLast(s)
	}
	fmt.Println(linkList1)
	fmt.Println(IsPalindrome(linkList1.dummyHead.next))

	linkList2 := Constructor()
	for _, s := range testStr2 {
		linkList2.AddLast(s)
	}
	fmt.Println(linkList2)
	fmt.Println(IsPalindrome(linkList2.dummyHead.next))
}
