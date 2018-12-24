package AVLSet

import "Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/src/AVLTree"

type AVLSet struct {
	avl AVLTree.AVLTree
}

func Constructor() *AVLSet {
	return &AVLSet{}
}

func (this *AVLSet) Add(k interface{}) {
	this.avl.Add(k, nil)
}

func (this *AVLSet) Remove(k interface{}) {
	this.avl.Remove(k)
}

func (this *AVLSet) Contains(k interface{}) bool {
	return this.avl.Contains(k)
}

func (this *AVLSet) GetSize() int {
	return this.avl.GetSize()
}

func (this *AVLSet) IsEmpty() bool {
	return this.avl.IsEmpty()
}
