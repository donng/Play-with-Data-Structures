package main

import (
	"Play-with-Data-Structures/Utils/FileOperation"
	"Play-with-Data-Structures/Utils/Interfaces"
	"bytes"
	"fmt"
	"path/filepath"
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

// 生成node节点，默认为红色
func generateNode(key interface{}, value interface{}) *Node {
	return &Node{key: key, value: value, color: RED}
}

// 返回以node为根节点的二分搜索树中，key所在的节点
func (rt *RBTree) getNode(n *Node, key interface{}) *Node {
	// 未找到等于 key 的节点
	if n == nil {
		return nil
	}

	if Interfaces.Compare(key, n.key) == 0 {
		return n
	} else if Interfaces.Compare(key, n.key) < 0 {
		return rt.getNode(n.left, key)
	} else {
		return rt.getNode(n.right, key)
	}
}

// 判断节点node的颜色
func isRed(n *Node) bool {
	if n == nil {
		return BLACK
	}
	return n.color
}

//   node                     x
//  /   \     左旋转         /  \
// T1   x   --------->   node   T3
//     / \              /   \
//    T2 T3            T1   T2
func leftRotate(n *Node) *Node {
	x := n.right

	// 左旋转
	n.right = x.left
	x.left = n

	x.color = n.color
	n.color = RED

	return x
}

//     node                   x
//    /   \     右旋转       /  \
//   x    T2   ------->   y   node
//  / \                       /  \
// y  T1                     T1  T2
func rightRotate(n *Node) *Node {
	x := n.left

	// 右旋转
	n.left = x.right
	x.right = n

	x.color = n.color
	n.color = RED

	return x
}

// 颜色翻转
func (node *Node) flipColors() {
	node.color = RED
	node.left.color = BLACK
	node.right.color = BLACK
}

// 向红黑树中添加新的元素(key, value)
func (rt *RBTree) Add(key interface{}, val interface{}) {
	rt.root = rt.add(rt.root, key, val)
	rt.root.color = BLACK // 最终根节点为黑色节点
}

// 向以node为根的红黑树中插入元素(key, value)，递归算法
// 返回插入新节点后红黑树的根
func (rt *RBTree) add(n *Node, key interface{}, val interface{}) *Node {
	if n == nil {
		rt.size++
		return generateNode(key, val)
	}

	if Interfaces.Compare(key, n.key) < 0 {
		n.left = rt.add(n.left, key, val)
	} else if Interfaces.Compare(key, n.key) > 0 {
		n.right = rt.add(n.right, key, val)
	} else {
		n.value = val
	}

	if isRed(n.right) && !isRed(n.left) {
		n = leftRotate(n)
	}
	if isRed(n.left) && isRed(n.left.left) {
		n = rightRotate(n)
	}
	if isRed(n.left) && isRed(n.right) {
		n.flipColors()
	}

	return n
}

// 从二分搜索树中删除键为key的节点
func (rt *RBTree) Remove(key interface{}) interface{} {
	n := rt.getNode(rt.root, key)
	if n != nil {
		rt.root = rt.remove(rt.root, key)
		return n.value
	}

	return nil
}

func (rt *RBTree) remove(n *Node, key interface{}) *Node {
	if n == nil {
		return nil
	}

	if Interfaces.Compare(key, n.key) < 0 {
		n.left = rt.remove(n.left, key)
		return n
	} else if Interfaces.Compare(key, n.key) > 0 {
		n.right = rt.remove(n.right, key)
		return n
	} else {
		// 待删除节点左子树为空的情况
		if n.left == nil {
			rightNode := n.right
			n.right = nil
			rt.size--
			return rightNode
		}
		// 待删除节点右子树为空的情况
		if n.right == nil {
			leftNode := n.left
			n.left = nil
			rt.size--
			return leftNode
		}
		// 待删除节点左右子树均不为空的情况

		// 找到比待删除节点大的最小节点, 即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		successor := rt.minimum(n.right)
		successor.right = rt.removeMin(n.right)
		successor.left = n.left

		n.left, n.right = nil, nil

		return successor
	}
}

// 返回以node为根的二分搜索树的最小值所在的节点
func (rt *RBTree) minimum(n *Node) *Node {
	if n.left == nil {
		return n
	}
	return rt.minimum(n.left)
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (rt *RBTree) removeMin(n *Node) *Node {
	if n.left == nil {
		rightNode := n.right
		n.right = nil
		rt.size--

		return rightNode
	}

	n.left = rt.removeMin(n.left)
	return n
}

func (rt *RBTree) Contains(key interface{}) bool {
	return rt.getNode(rt.root, key) != nil
}

func (rt *RBTree) Get(key interface{}) interface{} {
	n := rt.getNode(rt.root, key)
	if n == nil {
		return nil
	} else {
		return n.value
	}
}

func (rt *RBTree) Set(key interface{}, val interface{}) {
	n := rt.getNode(rt.root, key)
	if n == nil {
		panic(fmt.Sprintf("%v, doesn't exist", key))
	}

	n.value = val
}

func (rt *RBTree) GetSize() int {
	return rt.size
}

func (rt *RBTree) IsEmpty() bool {
	return rt.size == 0
}

func (rt *RBTree) String() string {
	var buffer bytes.Buffer
	generateBSTSting(rt.root, 0, &buffer)
	return buffer.String()
}

// 生成以node为根节点，深度为depth的描述二叉树的字符串
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

func main() {
	filename, _ := filepath.Abs("13-Red-Black-Tree/07-Adding-Elements-in-Red-Black-Tree/pride-and-prejudice.txt")

	RbTree := Constructor()

	words := FileOperation.ReadFile(filename)
	for _, word := range words {
		if RbTree.Contains(word) {
			RbTree.Set(word, RbTree.Get(word).(int)+1)
		} else {
			RbTree.Add(word, 1)
		}
	}

	fmt.Println("Total different words: ", RbTree.GetSize())
	fmt.Println("Frequency of PRIDE:", RbTree.Get("pride"))
	fmt.Println("Frequency of PREJUDICE: ", RbTree.Get("prejudice"))
}
