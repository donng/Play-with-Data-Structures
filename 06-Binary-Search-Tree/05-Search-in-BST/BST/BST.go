package BST

import "github.com/donng/Play-with-Data-Structures/utils"

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

func New() *BST {
	return &BST{}
}

func (b *BST) GetSize() interface{} {
	return b.size
}

func (b *BST) IsEmpty() bool {
	return b.size == 0
}

// 向二分搜索树中添加新的元素 e
func (b *BST) Add(e interface{}) {
	b.root = b.add(b.root, e)
}

// 向以 Node 为跟的二分搜索树中插入元素 e，递归算法
// 返回插入新节点后二分搜索树的根
func (b *BST) add(n *Node, e interface{}) *Node {
	if n == nil {
		b.size++
		return generateNode(e)
	}

	// 递归调用
	if utils.Compare(e, n.e) < 0 {
		n.left = b.add(n.left, e)
	} else if utils.Compare(e, n.e) > 0 {
		n.right = b.add(n.right, e)
	}
	return n
}

// 看二分搜索树中是否包含元素 e
func (b *BST) Contains(e interface{}) bool {
	return contains(b.root, e)
}

// 看以 Node 为根的二分搜索树是否包含元素 e，递归算法
func contains(n *Node, e interface{}) bool {
	if n == nil {
		return false
	}

	if utils.Compare(e, n.e) == 0 {
		return true
	} else if utils.Compare(e, n.e) < 0 {
		return contains(n.left, e)
	} else {
		return contains(n.right, e)
	}
}
