package AVLTree

import (
	"Play-with-Data-Structures/Utils/FileOperation"
	"Play-with-Data-Structures/Utils/Interfaces"
	"bytes"
	"fmt"
	"math"
	"path/filepath"
)

type Node struct {
	key         interface{}
	val         interface{}
	left, right *Node
	height      int
}

type AVLTree struct {
	root *Node
	size int
}

// 生成 Node 节点
func generateNode(k interface{}, v interface{}) *Node {
	return &Node{key: k, val: v, height: 1}
}

func Constructor() *AVLTree {
	return &AVLTree{}
}

// 判断该二叉树是否是一颗二分搜索树
func (this *AVLTree) IsBST() bool {
	var keys []interface{}
	inOrder(this.root, keys)

	for i := 1; i < len(keys); i++ {
		if Interfaces.Compare(keys[i-1], keys[i]) == 1 {
			return false
		}
	}
	return true
}

func inOrder(n *Node, keys []interface{}) {
	if n == nil {
		return
	}

	inOrder(n.left, keys)
	keys = append(keys, n.key)
	inOrder(n.right, keys)
}

// 判断该二叉树是否是一棵平衡二叉树
func (this *AVLTree) IsBalanced() bool {
	return this.isBalanced(this.root)
}

// 判断以Node为根的二叉树是否是一棵平衡二叉树，递归算法
func (this *AVLTree) isBalanced(n *Node) bool {
	if n == nil {
		return true
	}

	balanceFactor := this.getBalanceFactor(n)
	if math.Abs(float64(balanceFactor)) > 1 {
		return false
	}
	return this.isBalanced(n.left) && this.isBalanced(n.right)
}

// 返回以 Node 为根节点的二分搜索树中，key所在的节点
func (this *AVLTree) getNode(Node *Node, key interface{}) *Node {
	// 未找到等于 key 的节点
	if Node == nil {
		return nil
	}

	if Interfaces.Compare(key, Node.key) == 0 {
		return Node
	} else if Interfaces.Compare(key, Node.key) == -1 {
		return this.getNode(Node.left, key)
	} else {
		return this.getNode(Node.right, key)
	}
}

// 获得节点 Node 的高度
func (this *AVLTree) getHeight(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

// 获得节点 Node 的平衡因子
func (this *AVLTree) getBalanceFactor(n *Node) int {
	if n == nil {
		return 0
	}
	return this.getHeight(n.left) - this.getHeight(n.right)
}

// 对节点y进行向右旋转操作，返回旋转后新的根节点x
//        y                              x
//       / \                           /   \
//      x   T4     向右旋转 (y)        z     y
//     / \       - - - - - - - ->    / \   / \
//    z   T3                       T1  T2 T3 T4
//   / \
// T1   T2
func (this *AVLTree) rightRotate(y *Node) *Node {
	x := y.left
	T3 := x.right

	// 向右旋转过程
	x.right = y
	y.left = T3

	// 更新 height
	y.height = int(math.Max(float64(this.getHeight(y.left)), float64(this.getHeight(y.right)))) + 1
	x.height = int(math.Max(float64(this.getHeight(x.left)), float64(this.getHeight(x.right)))) + 1

	return x
}

// 对节点y进行向左旋转操作，返回旋转后新的根节点x
//    y                             x
//  /  \                          /   \
// T1   x      向左旋转 (y)       y     z
//     / \   - - - - - - - ->   / \   / \
//   T2  z                     T1 T2 T3 T4
//      / \
//     T3 T4
func (this *AVLTree) leftRotate(y *Node) *Node {
	x := y.right
	T2 := x.left

	// 向左旋转过程
	x.left = y
	y.right = T2

	// 更新 height
	y.height = int(math.Max(float64(this.getHeight(y.left)), float64(this.getHeight(y.right)))) + 1
	x.height = int(math.Max(float64(this.getHeight(x.left)), float64(this.getHeight(x.right)))) + 1

	return x
}

// 向二分搜索树中添加新的元素(key, value)
func (this *AVLTree) Add(key interface{}, val interface{}) {
	this.root = this.add(this.root, key, val)
}

// 向以node为根的二分搜索树中插入元素(key, value)，递归算法
// 返回插入新节点后二分搜索树的根
func (this *AVLTree) add(n *Node, key interface{}, val interface{}) *Node {
	if n == nil {
		this.size++
		return generateNode(key, val)
	}

	if Interfaces.Compare(key, n.key) < 0 {
		n.left = this.add(n.left, key, val)
	} else if Interfaces.Compare(key, n.key) > 0 {
		n.right = this.add(n.right, key, val)
	} else {
		n.val = val
	}

	// 更新 height
	n.height = 1 + int(math.Max(float64(this.getHeight(n.left)), float64(this.getHeight(n.right))))
	// 计算平衡因子
	balanceFactor := this.getBalanceFactor(n)
	//if math.Abs(float64(balanceFactor)) > 1 {
	//	fmt.Println("unbalanced: ", balanceFactor)
	//}
	// 平衡维护
	// LL
	if balanceFactor > 1 && this.getBalanceFactor(n.left) >= 0 {
		return this.rightRotate(n)
	}
	// RR
	if balanceFactor < -1 && this.getBalanceFactor(n.right) <= 0 {
		return this.leftRotate(n)
	}
	// LR
	if balanceFactor > 1 && this.getBalanceFactor(n.left) < 0 {
		n.left = this.leftRotate(n.left)
		return this.rightRotate(n)
	}
	// RL
	if balanceFactor < -1 && this.getBalanceFactor(n.right) > 0 {
		n.right = this.rightRotate(n.right)
		return this.leftRotate(n)
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

func (this *AVLTree) remove(n *Node, key interface{}) *Node {
	if n == nil {
		return nil
	}

	if Interfaces.Compare(key, n.key) == -1 {
		n.left = this.remove(n.left, key)
		return n
	} else if Interfaces.Compare(key, n.key) == 1 {
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
func (this *AVLTree) minimum(n *Node) *Node {
	if n.left == nil {
		return n
	}
	return this.minimum(n.left)
}

// 删除掉以node为根的二分搜索树中的最小节点
// 返回删除节点后新的二分搜索树的根
func (this *AVLTree) removeMin(n *Node) *Node {
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

// 生成以 Node 为根节点，深度为 depth 的描述二叉树的字符串
func generateBSTSting(Node *Node, depth int, buffer *bytes.Buffer) {
	if Node == nil {
		buffer.WriteString(generateDepthString(depth) + "nil\n")
		return
	}

	buffer.WriteString(generateDepthString(depth) + fmt.Sprintf("%s", Node.key) + "\n")
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

func main() {
	filename, _ := filepath.Abs("12-AVL-Tree/06-LR-and-RL/pride-and-prejudice.txt")

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

	fmt.Println("is BST:", AVL.IsBST())
	fmt.Println("is Balanced", AVL.IsBalanced())
}
