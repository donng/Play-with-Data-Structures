package BST

import (
	"Play-with-Data-Structures/Utils/Interfaces"
)

type Node struct {
	e     interface{}
	left  *Node
	right *Node
}

func generateNode(e interface{}) *Node {
	return &Node{e: e}
}

type BST struct {
	root *Node
	size int
}

func Constructor() *BST {
	return &BST{}
}

func (this *BST) GetSize() int {
	return this.size
}

func (this *BST) IsEmpty() bool {
	return this.size == 0
}

// 向二分搜索树中添加新的元素e
func (this *BST) Add(e interface{}) {
	// tree 为空，设置根节点
	if this.root == nil {
		this.root = generateNode(e)
		this.size++
	} else {
		this.add(this.root, e)
	}
}

// 向以node为根的二分搜索树中插入元素e，递归算法
func (this *BST) add(n *Node, e interface{}) {
	// 不处理重复数据的节点
	if e == n.e {
		return
		// 左子树递归终止条件
	} else if Interfaces.Compare(e, n.e) < 0 && n.left == nil {
		n.left = generateNode(e)
		this.size++
		return
		// 右子树递归终止条件
	} else if Interfaces.Compare(e, n.e) > 0 && n.right == nil {
		n.right = generateNode(e)
		this.size++
		return
	}

	// 递归调用
	if Interfaces.Compare(e, n.e) < 0 {
		this.add(n.left, e)
	} else {
		this.add(n.right, e)
	}
}
