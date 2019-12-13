package SegmentTree

type SegmentTree struct {
	tree []interface{}
	data []interface{}
}

func Constructor(arr []interface{}) *SegmentTree {
	segmentTree := &SegmentTree{
		tree: make([]interface{}, len(arr)*4),
		data: make([]interface{}, len(arr)),
	}

	for i := 0; i < len(arr); i++ {
		segmentTree.data[i] = arr[i]
	}

	return segmentTree
}

func (st *SegmentTree) GetSize() int {
	return len(st.data)
}

func (st *SegmentTree) Get(index int) interface{} {
	if index < 0 || index >= len(st.data) {
		panic("Index is illegal.")
	}

	return st.data[index]
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
func leftChild(index int) int {
	return index*2 + 1
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
func rightChild(index int) int {
	return index*2 + 2
}
