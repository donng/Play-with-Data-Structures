package SegmentTree

import (
	"bytes"
	"fmt"
)

type SegmentTree struct {
	tree   []interface{}
	data   []interface{}
	merger func(interface{}, interface{}) interface{}
}

func GetSegmentTree(arr []interface{}, merger func(interface{}, interface{}) interface{}) *SegmentTree {
	segmentTree := &SegmentTree{
		tree:   make([]interface{}, len(arr)*4),
		data:   make([]interface{}, len(arr)),
		merger: merger,
	}

	for i := 0; i < len(arr); i++ {
		segmentTree.data[i] = arr[i]
	}
	segmentTree.buildSegmentTree(0, 0, len(arr)-1)

	return segmentTree
}

// 在treeIndex的位置创建表示区间[l...r]的线段树
func (t *SegmentTree) buildSegmentTree(treeIndex int, l int, r int) {
	if l == r {
		t.tree[treeIndex] = t.data[l]
		return
	}
	leftTreeIndex := leftChild(treeIndex)
	rightTreeIndex := rightChild(treeIndex)

	mid := l + (r-l)/2
	t.buildSegmentTree(leftTreeIndex, l, mid)
	t.buildSegmentTree(rightTreeIndex, mid+1, r)

	t.tree[treeIndex] = t.merger(t.tree[leftTreeIndex], t.tree[rightTreeIndex])
}

func (t *SegmentTree) GetSize() int {
	return len(t.data)
}

func (t *SegmentTree) Get(index int) interface{} {
	if index < 0 || index >= len(t.data) {
		panic("Index is illegal.")
	}

	return t.data[index]
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
func leftChild(index int) int {
	return index*2 + 1
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
func rightChild(index int) int {
	return index*2 + 2
}

func (t *SegmentTree) String() string {
	buffer := bytes.Buffer{}

	buffer.WriteString("[")
	for i := 0; i < len(t.tree); i ++ {
		if t.tree[i] != nil {
			buffer.WriteString(fmt.Sprint(t.tree[i]))
		} else {
			buffer.WriteString("nil")
		}

		if i != len(t.tree)-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
