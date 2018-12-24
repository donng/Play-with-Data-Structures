package LinkedListSet

import (
	"Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/src/LinkedList"
)

type LinkedListSet struct {
	LinkedList *LinkedList.LinkedList
}

func Constructor() *LinkedListSet {
	list := LinkedList.GetLinkedList()
	return &LinkedListSet{list}
}

func (s *LinkedListSet) Add(e interface{}) {
	if !s.LinkedList.Contains(e) {
		s.LinkedList.AddFirst(e)
	}
}

func (s *LinkedListSet) Remove(e interface{}) {
	s.LinkedList.RemoveElement(e)
}

func (s *LinkedListSet) Contains(e interface{}) bool {
	return s.LinkedList.Contains(e)
}

func (s *LinkedListSet) GetSize() int {
	return s.LinkedList.GetSize()
}

func (s *LinkedListSet) IsEmpty() bool {
	return s.LinkedList.IsEmpty()
}
