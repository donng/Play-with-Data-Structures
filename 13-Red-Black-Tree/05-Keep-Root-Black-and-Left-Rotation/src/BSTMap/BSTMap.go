package BSTMap

import (
	"Play-with-Data-Structures/Utils/Interfaces"
	"bytes"
	"fmt"
)

type Node struct {
	key   interface{}
	val   interface{}
	left  *Node
	right *Node
}

type BSTMap struct {
	root *Node
	size int
}

func Constructor() *BSTMap {
	return &BSTMap{}
}

// 返回以node为根节点的二分搜索树中，key所在的节点
func (bm *BSTMap) getNode(n *Node, key interface{}) *Node {
	// 未找到等于 key 的节点
	if n == nil {
		return nil
	}

	if Interfaces.Compare(key, n.key) == 0 {
		return n
	} else if Interfaces.Compare(key, n.key) < 0 {
		return bm.getNode(n.left, key)
	} else {
		return bm.getNode(n.right, key)
	}
}

// 向二分搜索树中添加新的元素(key, value)
func (bm *BSTMap) Add(key interface{}, val interface{}) {
	bm.root = bm.add(bm.root, key, val)
}

// 向以node为根的二分搜索树中插入元素(key, value)，递归算法
// 返回插入新节点后二分搜索树的根
func (bm *BSTMap) add(n *Node, key interface{}, val interface{}) *Node {
	if n == nil {
		bm.size++
		return &Node{
			key: key,
			val: val,
		}
	}

	if Interfaces.Compare(key, n.key) < 0 {
		n.left = bm.add(n.left, key, val)
	} else if Interfaces.Compare(key, n.key) > 0 {
		n.right = bm.add(n.right, key, val)
	} else {
		n.val = val
	}

	return n
}

// 从二分搜索树中删除键为key的节点
func (bm *BSTMap) Remove(key interface{}) interface{} {
	n := bm.getNode(bm.root, key)
	if n != nil {
		bm.root = bm.remove(bm.root, key)
		return n.val
	}

	return nil
}

func (bm *BSTMap) remove(n *Node, key interface{}) *Node {
	if n == nil {
		return nil
	}

	if Interfaces.Compare(key, n.key) < 0 {
		n.left = bm.remove(n.left, key)
		return n
	} else if Interfaces.Compare(key, n.key) > 0 {
		n.right = bm.remove(n.right, key)
		return n
	} else {
		// 待删除节点左子树为空的情况
		if n.left == nil {
			rightNode := n.right
			n.right = nil
			bm.size--
			return rightNode
		}
		// 待删除节点右子树为空的情况
		if n.right == nil {
			leftNode := n.left
			n.left = nil
			bm.size--
			return leftNode
		}
		// 待删除节点左右子树均不为空的情况

		// 找到比待删除节点大的最小节点, 即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		successor := bm.minimum(n.right)
		successor.right = bm.removeMin(n.right)
		successor.left = n.left

		n.left, n.right = nil, nil

		return successor
	}
}

// 返回以node为根的二分搜索树的最小值所在的节点
func (bm *BSTMap) minimum(n *Node) *Node {
	if n.left == nil {
		return n
	}
	return bm.minimum(n.left)
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (bm *BSTMap) removeMin(n *Node) *Node {
	if n.left == nil {
		rightNode := n.right
		n.right = nil
		bm.size--

		return rightNode
	}

	n.left = bm.removeMin(n.left)
	return n
}

func (bm *BSTMap) Contains(key interface{}) bool {
	return bm.getNode(bm.root, key) != nil
}

func (bm *BSTMap) Get(key interface{}) interface{} {
	n := bm.getNode(bm.root, key)
	if n == nil {
		return nil
	} else {
		return n.val
	}
}

func (bm *BSTMap) Set(key interface{}, val interface{}) {
	n := bm.getNode(bm.root, key)
	if n == nil {
		panic(fmt.Sprintf("%v, doesn't exist", key))
	}

	n.val = val
}

func (bm *BSTMap) GetSize() int {
	return bm.size
}

func (bm *BSTMap) IsEmpty() bool {
	return bm.size == 0
}

func (bm *BSTMap) String() string {
	var buffer bytes.Buffer
	generateBSTSting(bm.root, 0, &buffer)
	return buffer.String()
}

// 生成以 Node 为根节点，深度为 depth 的描述二叉树的字符串
func generateBSTSting(Node *Node, depth int, buffer *bytes.Buffer) {
	if Node == nil {
		buffer.WriteString(generateDepthString(depth) + "nil\n")
		return
	}

	buffer.WriteString(generateDepthString(depth) + fmt.Sprint(Node.val) + "\n")
	generateBSTSting(Node.left, depth+1, buffer)
	generateBSTSting(Node.right, depth+1, buffer)
}

func generateDepthString(depth int) string {
	var buffer bytes.Buffer
	for i := 0; i < depth; i++ {
		buffer.WriteString("--")
	}
	return buffer.String()
}
