package linkedlistmap

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

func New() *LinkedListMap {
	return &LinkedListMap{
		dummyHead: &Node{},
	}
}

func (lm *LinkedListMap) GetSize() int {
	return lm.size
}

func (lm *LinkedListMap) IsEmpty() bool {
	return lm.size == 0
}

func (lm *LinkedListMap) getNode(key interface{}) *Node {
	prev := lm.dummyHead.next
	for prev != nil {
		if prev.key == key {
			return prev
		}
		prev = prev.next
	}

	return nil
}

func (lm *LinkedListMap) Add(key interface{}, val interface{}) {
	n := lm.getNode(key)

	if n == nil {
		newNode := &Node{
			key:  key,
			val:  val,
			next: lm.dummyHead.next,
		}
		lm.dummyHead.next = newNode
		lm.size++
	} else {
		n.val = val
	}
}

func (lm *LinkedListMap) Get(key interface{}) interface{} {
	n := lm.getNode(key)
	if n == nil {
		return nil
	} else {
		return n.val
	}
}

func (lm *LinkedListMap) Set(key interface{}, val interface{}) {
	Node := lm.getNode(key)
	if Node == nil {
		panic(fmt.Sprintf("%v, doesn't exist", key))
	}

	Node.val = val
}

func (lm *LinkedListMap) Remove(key interface{}) interface{} {
	prev := lm.dummyHead
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

func (lm *LinkedListMap) Contains(key interface{}) bool {
	return lm.getNode(key) != nil
}
