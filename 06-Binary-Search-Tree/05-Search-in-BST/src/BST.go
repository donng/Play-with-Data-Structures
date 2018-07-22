package src

type node struct {
	e     int
	left  *node
	right *node
}

type BST struct {
	root *node
	size int
}

func (t *BST) GetSize() int {
	return t.size
}

func (t *BST) IsEmpty() bool {
	return t.size == 0
}

// 向二分搜索树中添加新的元素 e
func (t *BST) Add(e int) {
	add(t.root, e)
}

// 向以 node 为跟的二分搜索树中插入元素 e，递归算法
// 返回插入新节点后二分搜索树的根
func (t *BST) add(node *node, e int) *node {
	if node == nil {
		t.size++
		return &node{e: e}
	}

	// 递归调用
	if e < node.e {
		node.left = t.add(node.left, e)
	} else if e > node.e {
		node.right = t.add(node.right, e)
	}
	return node
}

// 看二分搜索树中是否包含元素 e
func (t *BST) Contains(e int) bool {
	return contains(t.root, e)
}

// 看以 node 为根的二分搜索树是否包含元素 e，递归算法
func contains(node *node, e int) bool {
	if node == nil {
		return false
	}

	if e == node.e {
		return true
	} else if e < node.e {
		return contains(node.left, e)
	} else {
		return contains(node.right, e)
	}

}