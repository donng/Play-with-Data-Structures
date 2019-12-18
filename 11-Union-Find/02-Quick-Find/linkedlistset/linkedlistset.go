package linkedlistset

import "github.com/donng/Play-with-Data-Structures/11-Union-Find/02-Quick-Find/linkedlist"

type LinkedListSet struct {
	LinkedList *linkedlist.LinkedList
}

func Constructor() *LinkedListSet {
	return &LinkedListSet{
		LinkedList: linkedlist.New(),
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
