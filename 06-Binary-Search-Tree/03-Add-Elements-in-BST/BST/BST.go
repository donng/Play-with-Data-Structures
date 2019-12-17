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

func (b *BST) GetSize() int {
	return b.size
}

func (b *BST) IsEmpty() bool {
	return b.size == 0
}

// 向二分搜索树中添加新的元素e
func (b *BST) Add(e interface{}) {
	// tree 为空，设置根节点
	if b.root == nil {
		b.root = generateNode(e)
		b.size++
	} else {
		b.add(b.root, e)
	}
}

// 向以node为根的二分搜索树中插入元素e，递归算法
func (b *BST) add(n *Node, e interface{}) {
	// 不处理重复数据的节点
	if e == n.e {
		return
		// 左子树递归终止条件
	} else if utils.Compare(e, n.e) < 0 && n.left == nil {
		n.left = generateNode(e)
		b.size++
		return
		// 右子树递归终止条件
	} else if utils.Compare(e, n.e) > 0 && n.right == nil {
		n.right = generateNode(e)
		b.size++
		return
	}

	// 递归调用
	if utils.Compare(e, n.e) < 0 {
		b.add(n.left, e)
	} else {
		b.add(n.right, e)
	}
}
