package main

import (
	"fmt"
	"github.com/donng/Play-with-Data-Structures/utils"
)

type Node struct {
	e     interface{}
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
	size int
}

func New() *Tree {
	return &Tree{}
}

func (t *Tree) Add(e interface{}) {
	t.root = t.add(t.root, e)
}

func (t *Tree) add(n *Node, e interface{}) *Node {
	// 递归终止条件
	if n == nil {
		t.size++
		return &Node{e: e}
	}

	if utils.Compare(e, n.e) < 0 {
		n.left = t.add(n.left, e)
	} else if utils.Compare(e, n.e) > 0 {
		n.right = t.add(n.right, e)
	}

	return n
}

func (t *Tree) Contains(e interface{}) bool {
	return contains(t.root, e)
}

func contains(n *Node, e interface{}) bool {
	if n == nil {
		return false
	}

	if utils.Compare(e, n.e) < 0 {
		return contains(n.left, e)
	} else if utils.Compare(e, n.e) > 0 {
		return contains(n.right, e)
	} else {
		return true
	}
}

func (t *Tree) PreOrder() {
	preOrder(t.root)
}

func preOrder(n *Node) {
	if n == nil {
		return
	}

	fmt.Println(n.e)
	preOrder(n.left)
	preOrder(n.right)
}

func (t *Tree) InOrder() {
	inOrder(t.root)
}

func inOrder(n *Node) {
	if n == nil {
		return
	}

	inOrder(n.left)
	fmt.Println(n.e)
	inOrder(n.right)
}

// 后序遍历
func (t *Tree) PostOrder() {
	postOrder(t.root)
}

func postOrder(n *Node) {
	if n == nil {
		return
	}

	postOrder(n.left)
	postOrder(n.right)
	fmt.Println(n.e)
}

// 寻找二分搜索树的最小元素
func (t *Tree) Minimum() *Node {
	if t.size == 0 {
		panic("BST is empty!")
	}
	return minimum(t.root)
}

func minimum(n *Node) *Node {
	if n.left == nil {
		return n
	}
	return minimum(n.left)
}

// 寻找二分搜索树的最大元素
func (t *Tree) Maximum() *Node {
	if t.size == 0 {
		panic("BST is empty!")
	}
	return maximum(t.root)
}

func maximum(n *Node) *Node {
	if n.right == nil {
		return n
	}
	return maximum(n.right)
}

// 从二分搜索树中删除最小值所在的节点，返回最小值
func (t *Tree) RemoveMin() interface{} {
	retNode := t.Minimum()
	t.root = t.removeMin(t.root)
	return retNode
}

func (t *Tree) removeMin(n *Node) *Node {
	if n.left == nil {
		t.size--
		return n.right
	}
	n.left = t.removeMin(n.left)
	return n
}

func (t *Tree) RemoveMax() interface{} {
	retNode := t.Maximum()
	t.root = t.removeMax(t.root)
	return retNode
}

func (t *Tree) removeMax(n *Node) *Node {
	if n.right == nil {
		t.size--
		return n.left
	}
	n.right = t.removeMax(n.right)
	return n
}

func (t *Tree) RemoveElement(e interface{}) {
	t.root = t.remove(t.root, e)
}

func (t *Tree) remove(n *Node, e interface{}) *Node {
	// 递归到空节点，意味着树中不存在要删除的值
	if n == nil {
		return nil
	}

	if utils.Compare(e, n.e) < 0 {
		n.left = t.remove(n.left, e)
		return n
	} else if utils.Compare(e, n.e) > 0 {
		n.right = t.remove(n.right, e)
		return n
	} else {
		// 查找到要删除的节点 n.e == e
		if n.left == nil {
			rightNode := n.right
			n.right = nil
			t.size--
			return rightNode
		}
		if n.right == nil {
			leftNode := n.left
			n.left = nil
			t.size--
			return leftNode
		}

		// 左右子树都不为空时，取右子树的最小值代替删除的节点
		successor := minimum(n.right)
		successor.left = n.left
		successor.right = t.removeMin(n.right)
		n.left = nil
		n.right = nil
		return successor
	}
}
