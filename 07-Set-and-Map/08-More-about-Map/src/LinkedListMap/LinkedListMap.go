package LinkedListMap

import (
	"fmt"
)

type Node struct {
	key  interface{}
	val  interface{}
	next *Node
}

type LinkedListMap struct {
	dummyHead *Node
	size      int
}

func Constructor() *LinkedListMap {
	return &LinkedListMap{
		dummyHead: &Node{},
	}
}

func (this *LinkedListMap) GetSize() int {
	return this.size
}

func (this *LinkedListMap) IsEmpty() bool {
	return this.size == 0
}

func (this *LinkedListMap) getNode(key interface{}) *Node {
	prev := this.dummyHead.next
	for prev != nil {
		if prev.key == key {
			return prev
		}
		prev = prev.next
	}

	return nil
}

func (this *LinkedListMap) Add(key interface{}, val interface{}) {
	n := this.getNode(key)

	if n == nil {
		newNode := &Node{
			key:  key,
			val:  val,
			next: this.dummyHead.next,
		}
		this.dummyHead.next = newNode
		this.size++
	} else {
		n.val = val
	}
}

func (this *LinkedListMap) Get(key interface{}) interface{} {
	n := this.getNode(key)
	if n == nil {
		return nil
	} else {
		return n.val
	}
}

func (this *LinkedListMap) Set(key interface{}, val interface{}) {
	Node := this.getNode(key)
	if Node == nil {
		panic(fmt.Sprintf("%v, doesn't exist", key))
	}

	Node.val = val
}

func (this *LinkedListMap) Remove(key interface{}) interface{} {
	prev := this.dummyHead
	for prev.next != nil {
		if prev.next.key == key {
			delNode := prev.next
			prev.next = delNode.next
			delNode.next = nil

			return delNode.val
		}
		prev = prev.next
	}

	return nil
}

func (this *LinkedListMap) Contains(key interface{}) bool {
	return this.getNode(key) != nil
}
