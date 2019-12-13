package BSTSet

import (
	"Play-with-Data-Structures/10-Trie/04-Prefix-in-Trie/src/BST"
)

type BSTSet struct {
	BST *BST.BST
}

func Constructor() *BSTSet {
	return &BSTSet{
		BST: BST.Constructor(),
	}
}

func (bs *BSTSet) Add(e interface{}) {
	bs.BST.Add(e)
}

func (bs *BSTSet) Remove(e interface{}) {
	bs.BST.Remove(e)
}

func (bs *BSTSet) Contains(e interface{}) bool {
	return bs.BST.Contains(e)
}

func (bs *BSTSet) GetSize() int {
	return bs.BST.GetSize()
}

func (bs *BSTSet) IsEmpty() bool {
	return bs.BST.IsEmpty()
}
