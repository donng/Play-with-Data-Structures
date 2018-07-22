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
	t.add(t.root, e)
}

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
