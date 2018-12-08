package main

import (
	"bytes"
	"fmt"
	"math"
	"path/filepath"
	"Play-with-Data-Structures/12-AVL-Tree/02-Calculating-Balance-Factor/src/FileOperation"
	"reflect"
)

type node struct {
	key         interface{}
	val         interface{}
	left, right *node
	height      int
}

type AVLTree struct {
	root *node
	size int
}

// 生成 node 节点
func generateNode() *node {
	return &node{key: "", val: 0, height: 1}
}

func Constructor() *AVLTree {
	return &AVLTree{root: generateNode()}
}

// interface{} 类型的比较
func compare(a interface{}, b interface{}) int {
	aType := reflect.TypeOf(a).String()
	bType := reflect.TypeOf(b).String()
	if aType != bType {
		panic("cannot compare different type params")
	}

	switch a.(type) {
	case int:
		if a.(int) > b.(int) {
			return 1
		} else if a.(int) < b.(int) {
			return -1
		} else {
			return 0
		}
	case string:
		if a.(string) > b.(string) {
			return 1
		} else if a.(string) < b.(string) {
			return -1
		} else {
			return 0
		}
	case float64:
		if a.(float64) > b.(float64) {
			return 1
		} else if a.(float64) < b.(float64) {
			return -1
		} else {
			return 0
		}
	default:
		panic("unsupported type params")
	}
}

// 返回以 node 为根节点的二分搜索树中，key所在的节点
func (this *AVLTree) getNode(node *node, key interface{}) *node {
	// 未找到等于 key 的节点
	if node == nil {
		return nil
	}

	if compare(key, node.key) == 0 {
		return node
	} else if compare(key, node.key) == -1 {
		return this.getNode(node.left, key)
	} else {
		return this.getNode(node.right, key)
	}
}

// 获得节点 node 的高度
func (this *AVLTree) getHeight(n *node) int {
	if n == nil {
		return 0
	} else {
		return n.height
	}
}

// 获得节点 node 的平衡因子
func (this *AVLTree) getBalanceFactor(n *node) int {
	if n == nil {
		return 0
	} else {
		return this.getHeight(n.left) - this.getHeight(n.right)
	}
}

// 向二分搜索树中添加新的元素(key, value)
func (this *AVLTree) Add(key interface{}, val interface{}) {
	this.root = this.add(this.root, key, val)
}

// 向以node为根的二分搜索树中插入元素(key, value)，递归算法
// 返回插入新节点后二分搜索树的根
func (this *AVLTree) add(n *node, key interface{}, val interface{}) *node {
	if n == nil {
		this.size++
		return &node{
			key: key,
			val: val,
		}
	}

	if compare(key, n.key) == -1 {
		n.left = this.add(n.left, key, val)
	} else if compare(key, n.key) == 1 {
		n.right = this.add(n.right, key, val)
	} else {
		n.val = val
	}

	// 更新 height
	n.height = 1 + int(math.Max(float64(this.getHeight(n.left)), float64(this.getHeight(n.right))))
	// 计算平衡因子
	balanceFactor := this.getBalanceFactor(n)
	if math.Abs(float64(balanceFactor)) > 1 {
		fmt.Println("unbalanced: ", balanceFactor)
	}

	return n
}

// 从二分搜索树中删除键为 key 的节点
func (this *AVLTree) Remove(key interface{}) interface{} {
	n := this.getNode(this.root, key)
	if n != nil {
		this.root = this.remove(this.root, key)
		return n.val
	}

	return nil
}

func (this *AVLTree) remove(n *node, key interface{}) *node {
	if n == nil {
		return nil
	}

	if compare(key, n.key) == -1 {
		n.left = this.remove(n.left, key)
		return n
	} else if compare(key, n.key) == 1 {
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
func (this *AVLTree) minimum(n *node) *node {
	if n.left == nil {
		return n
	}
	return this.minimum(n.left)
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (this *AVLTree) removeMin(n *node) *node {
	if n.left == nil {
		rightNode := n.right
		n.right = nil
		this.size--

		return rightNode
	}

	n.left = this.removeMin(n.left)
	return n
}

func (this *AVLTree) Contains(key interface{}) bool {
	return this.getNode(this.root, key) != nil
}

func (this *AVLTree) Get(key interface{}) interface{} {
	n := this.getNode(this.root, key)
	if n == nil {
		return nil
	} else {
		return n.val
	}
}

func (this *AVLTree) Set(key interface{}, val interface{}) {
	n := this.getNode(this.root, key)
	if n == nil {
		panic(fmt.Sprintf("%v, doesn't exist", key))
	}

	n.val = val
}

func (this *AVLTree) GetSize() int {
	return this.size
}

func (this *AVLTree) IsEmpty() bool {
	return this.size == 0
}

func (this *AVLTree) String() string {
	var buffer bytes.Buffer
	generateBSTSting(this.root, 0, &buffer)
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

func main() {
	filename, _ := filepath.Abs("12-AVL-Tree/02-Calculating-Balance-Factor/pride-and-prejudice.txt")

	AVL := Constructor()

	words := FileOperation.ReadFile(filename)
	for _, word := range words {
		if AVL.Contains(word) {
			AVL.Set(word, AVL.Get(word).(int)+1)
		} else {
			AVL.Add(word, 1)
		}
	}

	fmt.Println("Total different words: ", AVL.GetSize())
	fmt.Println("Frequency of PRIDE:", AVL.Get("pride"))
	fmt.Println("Frequency of PREJUDICE: ", AVL.Get("prejudice"))
}
