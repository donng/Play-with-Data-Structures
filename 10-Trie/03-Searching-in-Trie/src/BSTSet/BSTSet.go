package BSTSet

import (
	"Play-with-Data-Structures/10-Trie/03-Searching-in-Trie/src/BST"
)

type BSTSet struct {
	BST *BST.BST
}

func Constructor() *BSTSet {
	return &BSTSet{
		BST: BST.Constructor(),
	}
}

func (this *BSTSet) Add(e interface{}) {
	this.BST.Add(e)
}

func (this *BSTSet) Remove(e interface{}) {
	this.BST.Remove(e)
}

func (this *BSTSet) Contains(e interface{}) bool {
	return this.BST.Contains(e)
}

func (this *BSTSet) GetSize() int {
	return this.BST.GetSize()
}

func (this *BSTSet) IsEmpty() bool {
	return this.BST.IsEmpty()
}
