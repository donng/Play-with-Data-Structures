package BSTSet

import "github.com/donng/Play-with-Data-Structures/10-Trie/03-Searching-in-Trie/BST"

type BSTSet struct {
	BST *BST.BST
}

func New() *BSTSet {
	return &BSTSet{
		BST: BST.New(),
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
