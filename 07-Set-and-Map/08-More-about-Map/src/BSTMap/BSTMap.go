package BSTMap

import (
	"fmt"
	"bytes"
)

type node struct {
	key   interface{}
	val   interface{}
	left  *node
	right *node
}

type BSTMap struct {
	root *node
	size int
}

func GetBSTMap() *BSTMap {
	return &BSTMap{}
}

// 返回以node为根节点的二分搜索树中，key所在的节点
func (m *BSTMap) getNode(node *node, key interface{}) *node {
	// 未找到等于 key 的节点
	if node == nil {
		return nil
	}

	if key.(string) == node.key.(string) {
		return node
	} else if key.(string) < node.key.(string) {
		return m.getNode(node.left, key)
	} else {
		return m.getNode(node.right, key)
	}
}

// 向二分搜索树中添加新的元素(key, value)
func (m *BSTMap) Add(key interface{}, val interface{}) {
	m.root = m.add(m.root, key, val)
}

// 向以node为根的二分搜索树中插入元素(key, value)，递归算法
// 返回插入新节点后二分搜索树的根
func (m *BSTMap) add(n *node, key interface{}, val interface{}) *node {
	if n == nil {
		m.size++
		return &node{
			key: key,
			val: val,
		}
	}

	if key.(string) < n.key.(string) {
		n.left = m.add(n.left, key, val)
	} else if key.(string) > n.key.(string) {
		n.right = m.add(n.right, key, val)
	} else {
		n.val = val
	}

	return n
}

// 从二分搜索树中删除键为key的节点
func (m *BSTMap) Remove(key interface{}) interface{} {
	n := m.getNode(m.root, key)
	if n != nil {
		m.root = m.remove(m.root, key)
		return n.val
	}

	return nil
}

func (m *BSTMap) remove(n *node, key interface{}) *node {
	if n == nil {
		return nil
	}

	if key.(string) < n.key.(string) {
		n.left = m.remove(n.left, key)
		return n
	} else if key.(string) > n.key.(string) {
		n.right = m.remove(n.right, key)
		return n
	} else {
		// 待删除节点左子树为空的情况
		if n.left == nil {
			rightNode := n.right
			n.right = nil
			m.size--
			return rightNode
		}
		// 待删除节点右子树为空的情况
		if n.right == nil {
			leftNode := n.left
			n.left = nil
			m.size--
			return leftNode
		}
		// 待删除节点左右子树均不为空的情况

		// 找到比待删除节点大的最小节点, 即待删除节点右子树的最小节点
		// 用这个节点顶替待删除节点的位置
		successor := m.minimum(n.right)
		successor.right = m.removeMin(n.right)
		successor.left = n.left

		n.left, n.right = nil, nil

		return successor
	}
}

// 返回以node为根的二分搜索树的最小值所在的节点
func (m *BSTMap) minimum(n *node) *node {
	if n.left == nil {
		return n
	}
	return m.minimum(n.left)
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (m *BSTMap) removeMin(n *node) *node {
	if n.left == nil {
		rightNode := n.right
		n.right = nil
		m.size--

		return rightNode
	}

	n.left = m.removeMin(n.left)
	return n
}

func (m *BSTMap) Contains(key interface{}) bool {
	return m.getNode(m.root, key) != nil
}

func (m *BSTMap) Get(key interface{}) interface{} {
	n := m.getNode(m.root, key)
	if n == nil {
		return nil
	} else {
		return n.val
	}
}

func (m *BSTMap) Set(key interface{}, val interface{}) {
	n := m.getNode(m.root, key)
	if n == nil {
		panic(fmt.Sprintf("%v, doesn't exist", key))
	}

	n.val = val
}

func (m *BSTMap) GetSize() int {
	return m.size
}

func (m *BSTMap) IsEmpty() bool {
	return m.size == 0
}

func (m *BSTMap) String() string {
	var buffer bytes.Buffer
	generateBSTSting(m.root, 0, &buffer)
	return buffer.String()
}

// 生成以 node 为根节点，深度为 depth 的描述二叉树的字符串
func generateBSTSting(node *node, depth int, buffer *bytes.Buffer) {
	if node == nil {
		buffer.WriteString(generateDepthString(depth) + "nil\n")
		return
	}

	buffer.WriteString(generateDepthString(depth) + fmt.Sprintf("%s", node.val) + "\n")
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
