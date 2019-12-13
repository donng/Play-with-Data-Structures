package LinkedListSet

import (
	"Play-with-Data-Structures/07-Set-and-Map/02-LinkedListSet/src/LinkedList"
)

type LinkedListSet struct {
	LinkedList *LinkedList.LinkedList
}

func Constructor() *LinkedListSet {
	return &LinkedListSet{
		LinkedList: LinkedList.Constructor(),
	}
}

func (ls *LinkedListSet) Add(e interface{}) {
	if !ls.LinkedList.Contains(e) {
		ls.LinkedList.AddFirst(e)
	}
}

func (ls *LinkedListSet) Remove(e interface{}) {
	ls.LinkedList.RemoveElement(e)
}

func (ls *LinkedListSet) Contains(e interface{}) bool {
	return ls.LinkedList.Contains(e)
}

func (ls *LinkedListSet) GetSize() int {
	return ls.LinkedList.GetSize()
}

func (ls *LinkedListSet) IsEmpty() bool {
	return ls.LinkedList.IsEmpty()
}
