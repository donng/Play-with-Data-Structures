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

func (t *BST) add(node *node, e int) {
	// 不处理重复数据的节点
	if e == node.e {
		return
		// 左子树递归终止条件
	} else if e < node.e && node.left == nil {
		node.left = &node{e: e}
		t.size++
		return
		// 右子树递归终止条件
	} else if e > node.e && node.right == nil {
		node.right = &node{e: e}
		t.size++
		return
	}

	// 递归调用
	if e < node.e {
		t.add(node.left, e)
	} else {
		t.add(node.right, e)
	}
}

