package AVLMap

import "Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/src/AVLTree"

type AVLMap struct {
	avl AVLTree.AVLTree
}

func Constructor() *AVLMap {
	return &AVLMap{}
}

func (this *AVLMap) Add(k interface{}, v interface{}) {
	this.avl.Add(k, v)
}

func (this *AVLMap) Remove(k interface{}) interface{} {
	return this.avl.Remove(k)
}

func (this *AVLMap) Contains(k interface{}) bool {
	return this.avl.Contains(k)
}

func (this *AVLMap) Get(k interface{}) interface{} {
	return this.avl.Get(k)
}

func (this *AVLMap) Set(k interface{}, v interface{}) {
	this.avl.Set(k, v)
}

func (this *AVLMap) GetSize() int {
	return this.avl.GetSize()
}

func (this *AVLMap) IsEmpty() bool {
	return this.avl.IsEmpty()
}
