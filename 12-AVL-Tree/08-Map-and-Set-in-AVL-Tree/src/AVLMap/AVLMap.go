package AVLMap

import "Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/src/AVLTree"

type AVLMap struct {
	avl AVLTree.AVLTree
}

func Constructor() *AVLMap {
	return &AVLMap{}
}

func (am *AVLMap) Add(k interface{}, v interface{}) {
	am.avl.Add(k, v)
}

func (am *AVLMap) Remove(k interface{}) interface{} {
	return am.avl.Remove(k)
}

func (am *AVLMap) Contains(k interface{}) bool {
	return am.avl.Contains(k)
}

func (am *AVLMap) Get(k interface{}) interface{} {
	return am.avl.Get(k)
}

func (am *AVLMap) Set(k interface{}, v interface{}) {
	am.avl.Set(k, v)
}

func (am *AVLMap) GetSize() int {
	return am.avl.GetSize()
}

func (am *AVLMap) IsEmpty() bool {
	return am.avl.IsEmpty()
}
