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

func (t *BST) Add(e int) {
	// tree 为空，设置根节点
	if t.root == nil {
		t.root = &node{e: e}
	} else {
		t.add(t.root, e)
	}
}

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

