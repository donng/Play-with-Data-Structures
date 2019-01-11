package BST

type Node struct {
	e     interface{}
	left  *Node
	right *Node
}

func generateNode(e interface{}) *Node {
	return &Node{e: e}
}

type BST struct {
	root *Node
	size int
}

func Constructor() *BST {
	return &BST{}
}

func (this *BST) GetSize() int {
	return this.size
}

func (this *BST) IsEmpty() bool {
	return this.size == 0
}
