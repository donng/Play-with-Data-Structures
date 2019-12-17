package BST

import (
	"bytes"
	"fmt"
	"github.com/donng/Play-with-Data-Structures/utils"
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

func New() *BST {
	return &BST{}
}

func (b *BST) GetSize() int {
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

// 二分搜索树的前序遍历
func (b *BST) PreOrder() {
	preOrder(b.root)
}

// 前序遍历以 Node 为根的二分搜索树，递归算法
func preOrder(n *Node) {
	if n == nil {
		return
	}

	fmt.Println(n.e)
	preOrder(n.left)
	preOrder(n.right)
}

// 二分搜索树的非递归前序遍历
//func (b *BST) PreOrderNR() {
//	// 使用之前我们自己实现的数组栈
//	stack := ArrayStack.Constructor(20)
//	stack.Push(b.root)
//
//	for !stack.IsEmpty() {
//		cur := stack.Pop().(*Node)
//		fmt.Println(cur.e)
//
//		if cur.right != nil {
//			stack.Push(cur.right)
//		}
//		if cur.left != nil {
//			stack.Push(cur.left)
//		}
//	}
//}

// 二分搜索树的中序遍历
func (b *BST) InOrder() {
	inOrder(b.root)
}

// 中序遍历以 Node 为根的二分搜索树，递归算法
func inOrder(n *Node) {
	if n == nil {
		return
	}

	inOrder(n.left)
	fmt.Println(n.e)
	inOrder(n.right)
}

// 二分搜索树的后序遍历
func (b *BST) PostOrder() {
	postOrder(b.root)
}

// 后序遍历以 Node 为根的二分搜索树，递归算法
func postOrder(n *Node) {
	if n == nil {
		return
	}

	postOrder(n.left)
	postOrder(n.right)
	fmt.Println(n.e)
}

// 二分搜索树的层序遍历
//func (b *BST) LevelOrder() {
//	// 使用我们之前实现的循环队列
//	queue := LoopQueue.Constructor(20)
//	queue.Enqueue(b.root)
//	for !queue.IsEmpty() {
//		cur := queue.Dequeue().(*Node)
//		fmt.Println(cur.e)
//
//		if cur.left != nil {
//			queue.Enqueue(cur.left)
//		}
//		if cur.right != nil {
//			queue.Enqueue(cur.right)
//		}
//	}
//}

// 寻找二分搜索树的最小元素
func (b *BST) Minimum() interface{} {
	if b.size == 0 {
		panic("BST is empty!")
	}
	return minimum(b.root).e
}

// 返回以 Node 为根的二分搜索树的最小值所在的节点
func minimum(n *Node) *Node {
	if n.left == nil {
		return n
	}
	return minimum(n.left)
}

// 寻找二分搜索树的最大元素
func (b *BST) Maximum() interface{} {
	if b.size == 0 {
		panic("BST is empty!")
	}
	return maximum(b.root).e
}

// 返回以 Node 为根的二分搜索树的最大值所在的节点
func maximum(n *Node) *Node {
	if n.right == nil {
		return n
	}
	return maximum(n.right)
}

// 从二分搜索树中删除最小值所在的节点，返回最小值
func (b *BST) RemoveMin() interface{} {
	// 获得最小值
	ret := b.Minimum()
	b.root = b.removeMin(b.root)
	return ret
}

// 删除以 Node 为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (b *BST) removeMin(n *Node) *Node {
	if n.left == nil {
		rightNode := n.right
		n.right = nil
		b.size--
		return rightNode
	}
	n.left = b.removeMin(n.left)
	return n
}

// 从二分搜索树中删除最小值所在的节点，返回最小值
func (b *BST) RemoveMax() interface{} {
	// 获得最小值
	ret := b.Maximum()
	b.root = b.removeMax(b.root)
	return ret
}

// 删除以 Node 为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (b *BST) removeMax(n *Node) *Node {
	if n.right == nil {
		leftNode := n.left
		n.left = nil
		b.size--
		return leftNode
	}
	n.right = b.removeMax(n.right)
	return n
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
