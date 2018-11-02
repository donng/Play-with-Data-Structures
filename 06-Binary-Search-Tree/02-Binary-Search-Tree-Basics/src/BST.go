package main

type node struct {
	e     interface{}
	left  *node
	right *node
}

type BST struct {
	root *node
	size int
}

func GetBST(e interface{}) *BST {
	bst := &BST{}
	bst.root = nil
	bst.size = 0

	return bst
}

func (t *BST) GetSize() int {
	return t.size
}

func (t *BST) IsEmpty() bool {
	return t.size == 0
}
