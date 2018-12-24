package BSTSet

import (
	"Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/src/BST"
)

type BSTSet struct {
	BST *BST.BST
}

func Constructor() *BSTSet {
	bst := BST.GetBST()
	return &BSTSet{bst}
}

func (s *BSTSet) Add(e interface{}) {
	s.BST.Add(e)
}

func (s *BSTSet) Remove(e interface{}) {
	s.BST.Remove(e)
}

func (s *BSTSet) Contains(e interface{}) bool {
	return s.BST.Contains(e)
}

func (s *BSTSet) GetSize() int {
	return s.BST.GetSize()
}

func (s *BSTSet) IsEmpty() bool {
	return s.BST.IsEmpty()
}
