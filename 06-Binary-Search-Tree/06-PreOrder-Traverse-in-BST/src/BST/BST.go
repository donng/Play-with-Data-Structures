package BST

import (
	"Play-with-Data-Structures/Utils/Interfaces"
	"bytes"
	"fmt"
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
	if Interfaces.Compare(e, n.e) < 0 {
		n.left = b.add(n.left, e)
	} else if Interfaces.Compare(e, n.e) > 0 {
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

	if Interfaces.Compare(e, n.e) == 0 {
		return true
	} else if Interfaces.Compare(e, n.e) < 0 {
		return contains(n.left, e)
	} else {
		return contains(n.right, e)
	}
}

func (b *BST) PreOrder() {
	preOrder(b.root)
}

func preOrder(n *Node) {
	if n == nil {
		return
	}

	fmt.Println(n.e)
	preOrder(n.left)
	preOrder(n.right)
}

func (b *BST) String() string {
	var buffer bytes.Buffer
	generateBSTSting(b.root, 0, &buffer)
	return buffer.String()
}

// 生成以 Node 为根节点，深度为 depth 的描述二叉树的字符串
func generateBSTSting(n *Node, depth int, buffer *bytes.Buffer) {
	if n == nil {
		buffer.WriteString(generateDepthString(depth) + "nil\n")
		return
	}

	buffer.WriteString(generateDepthString(depth) + fmt.Sprint(n.e) + "\n")
	generateBSTSting(n.left, depth+1, buffer)
	generateBSTSting(n.right, depth+1, buffer)
}

func generateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}
