package LinkedListSet

import (
	"Play-with-Data-Structures/11-Union-Find/02-Quick-Find/src/LinkedList"
)

type LinkedListSet struct {
	LinkedList *LinkedList.LinkedList
}

func Constructor() *LinkedListSet {
	return &LinkedListSet{
		LinkedList: LinkedList.Constructor(),
	}
}

func (this *LinkedListSet) Add(e interface{}) {
	if !this.LinkedList.Contains(e) {
		this.LinkedList.AddFirst(e)
	}
}

func (this *LinkedListSet) Remove(e interface{}) {
	this.LinkedList.RemoveElement(e)
}

func (this *LinkedListSet) Contains(e interface{}) bool {
	return this.LinkedList.Contains(e)
}

func (this *LinkedListSet) GetSize() int {
	return this.LinkedList.GetSize()
}

func (this *LinkedListSet) IsEmpty() bool {
	return this.LinkedList.IsEmpty()
}
