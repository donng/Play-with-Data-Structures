package BST

type node struct {
	e     int
	left  *node
	right *node
}

type BST struct {
	root *node
	size int
}

func GetBST(e interface{}) *BST {
	bst := &BST{}
	bst.root = nil
	bst.size = 0

	return bst
}

func (t *BST) GetSize() int {
	return t.size
}

func (t *BST) IsEmpty() bool {
	return t.size == 0
}

// 向二分搜索树中添加新的元素e
func (t *BST) Add(e int) {
	// tree 为空，设置根节点
	if t.root == nil {
		t.root = &node{e: e}
		t.size++
	} else {
		t.add(t.root, e)
	}
}

// 向以node为根的二分搜索树中插入元素e，递归算法
func (t *BST) add(n *node, e int) {
	// 不处理重复数据的节点
	if e == n.e {
		return
		// 左子树递归终止条件
	} else if e < n.e && n.left == nil {
		n.left = &node{e: e}
		t.size++
		return
		// 右子树递归终止条件
	} else if e > n.e && n.right == nil {
		n.right = &node{e: e}
		t.size++
		return
	}

	// 递归调用
	if e < n.e {
		t.add(n.left, e)
	} else {
		t.add(n.right, e)
	}
}
