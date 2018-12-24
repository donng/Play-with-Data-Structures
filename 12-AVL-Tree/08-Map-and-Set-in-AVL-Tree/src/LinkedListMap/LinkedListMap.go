package LinkedListMap

import (
	"fmt"
)

type node struct {
	key  interface{}
	val  interface{}
	next *node
}

type LinkedListMap struct {
	dummyHead *node
	size      int
}

func Constructor() *LinkedListMap {
	linkedListMap := &LinkedListMap{
		dummyHead: &node{},
	}

	return linkedListMap
}

func (l *LinkedListMap) GetSize() int {
	return l.size
}

func (l *LinkedListMap) IsEmpty() bool {
	return l.size == 0
}

func (l *LinkedListMap) getNode(key interface{}) *node {
	prev := l.dummyHead.next
	for prev != nil {
		if prev.key == key {
			return prev
		}
		prev = prev.next
	}

	return nil
}

func (l *LinkedListMap) Add(key interface{}, val interface{}) {
	n := l.getNode(key)

	if n == nil {
		newNode := &node{
			key:  key,
			val:  val,
			next: l.dummyHead.next,
		}
		l.dummyHead.next = newNode
		l.size++
	} else {
		n.val = val
	}
}

func (l *LinkedListMap) Get(key interface{}) interface{} {
	n := l.getNode(key)
	if n == nil {
		return nil
	} else {
		return n.val
	}
}

func (l *LinkedListMap) Set(key interface{}, val interface{}) {
	node := l.getNode(key)
	if node == nil {
		panic(fmt.Sprintf("%v, doesn't exist", key))
	}

	node.val = val
}

func (l *LinkedListMap) Remove(key interface{}) interface{} {
	prev := l.dummyHead
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

func (l *LinkedListMap) Contains(key interface{}) bool {
	return l.getNode(key) != nil
}
