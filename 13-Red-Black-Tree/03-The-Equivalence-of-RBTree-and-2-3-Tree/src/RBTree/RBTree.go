package RBTree

import (
	"Play-with-Data-Structures/Utils/Interfaces"
	"bytes"
	"fmt"
)

const RED = true
const BLACK = false

type Node struct {
	key, value  interface{}
	left, right *Node
	color       bool
}

type RBTree struct {
	root *Node
	size int
}

func Constructor() *RBTree {
	return &RBTree{}
}

// 生成 node 节点
func generateNode(key interface{}, value interface{}) *Node {
	return &Node{key: key, value: value, color: RED}
}

// 返回以node为根节点的二分搜索树中，key所在的节点
func (this *RBTree) getNode(n *Node, key interface{}) *Node {
	// 未找到等于 key 的节点
	if n == nil {
		return nil
	}

	if Interfaces.Compare(key, n.key) == 0 {
		return n
	} else if Interfaces.Compare(key, n.key) < 0 {
		return this.getNode(n.left, key)
	} else {
		return this.getNode(n.right, key)
	}
}

// 判断节点node的颜色
func isRed(n *Node) bool {
	if n == nil {
		return BLACK
	}
	return n.color
}

// 向二分搜索树中添加新的元素(key, value)
func (this *RBTree) Add(key interface{}, val interface{}) {
	this.root = this.add(this.root, key, val)
}

// 向以node为根的二分搜索树中插入元素(key, value)，递归算法
// 返回插入新节点后二分搜索树的根
func (this *RBTree) add(n *Node, key interface{}, val interface{}) *Node {
	if n == nil {
		this.size++
		return generateNode(key, val)
	}

	if Interfaces.Compare(key, n.key) < 0 {
		n.left = this.add(n.left, key, val)
	} else if Interfaces.Compare(key, n.key) > 0 {
		n.right = this.add(n.right, key, val)
	} else {
		n.value = val
	}

	return n
}

// 从二分搜索树中删除键为key的节点
func (this *RBTree) Remove(key interface{}) interface{} {
	n := this.getNode(this.root, key)
	if n != nil {
		this.root = this.remove(this.root, key)
		return n.value
	}

	return nil
}

func (this *RBTree) remove(n *Node, key interface{}) *Node {
	if n == nil {
		return nil
	}

	if Interfaces.Compare(key, n.key) < 0 {
		n.left = this.remove(n.left, key)
		return n
	} else if Interfaces.Compare(key, n.key) > 0 {
		n.right = this.remove(n.right, key)
		return n
	} else {
		// 待删除节点左子树为空的情况
		if n.left == nil {
			rightNode := n.right
			n.right = nil
			this.size--
			return rightNode
		}
		// 待删除节点右子树为空的情况
		if n.right == nil {
			leftNode := n.left
			n.left = nil
			this.size--
			return leftNode
		}
		// 待删除节点左右子树均不为空的情况

		// 找到比待删除节点大的最小节点, 即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		successor := this.minimum(n.right)
		successor.right = this.removeMin(n.right)
		successor.left = n.left

		n.left, n.right = nil, nil

		return successor
	}
}

// 返回以node为根的二分搜索树的最小值所在的节点
func (this *RBTree) minimum(n *Node) *Node {
	if n.left == nil {
		return n
	}
	return this.minimum(n.left)
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (this *RBTree) removeMin(n *Node) *Node {
	if n.left == nil {
		rightNode := n.right
		n.right = nil
		this.size--

		return rightNode
	}

	n.left = this.removeMin(n.left)
	return n
}

func (this *RBTree) Contains(key interface{}) bool {
	return this.getNode(this.root, key) != nil
}

func (this *RBTree) Get(key interface{}) interface{} {
	n := this.getNode(this.root, key)
	if n == nil {
		return nil
	} else {
		return n.value
	}
}

func (this *RBTree) Set(key interface{}, val interface{}) {
	n := this.getNode(this.root, key)
	if n == nil {
		panic(fmt.Sprintf("%v, doesn't exist", key))
	}

	n.value = val
}

func (this *RBTree) GetSize() int {
	return this.size
}

func (this *RBTree) IsEmpty() bool {
	return this.size == 0
}

func (this *RBTree) String() string {
	var buffer bytes.Buffer
	generateBSTSting(this.root, 0, &buffer)
	return buffer.String()
}

// 生成以 node 为根节点，深度为 depth 的描述二叉树的字符串
func generateBSTSting(node *Node, depth int, buffer *bytes.Buffer) {
	if node == nil {
		buffer.WriteString(generateDepthString(depth) + "nil\n")
		return
	}

	buffer.WriteString(generateDepthString(depth) + fmt.Sprintf("%s", node.value) + "\n")
	generateBSTSting(node.left, depth+1, buffer)
	generateBSTSting(node.right, depth+1, buffer)
}

func generateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}
