package AVLSet

import "github.com/donng/Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/AVLTree"

type AVLSet struct {
	avl AVLTree.AVLTree
}

func New() *AVLSet {
	return &AVLSet{}
}

func (as *AVLSet) Add(k interface{}) {
	as.avl.Add(k, nil)
}

func (as *AVLSet) Remove(k interface{}) {
	as.avl.Remove(k)
}

func (as *AVLSet) Contains(k interface{}) bool {
	return as.avl.Contains(k)
}

func (as *AVLSet) GetSize() int {
	return as.avl.GetSize()
}

func (as *AVLSet) IsEmpty() bool {
	return as.avl.IsEmpty()
}
