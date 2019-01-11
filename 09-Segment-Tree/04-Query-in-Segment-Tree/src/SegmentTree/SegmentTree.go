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

func Constructor(arr []interface{}, merger func(interface{}, interface{}) interface{}) *SegmentTree {
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
func (this *SegmentTree) buildSegmentTree(treeIndex int, l int, r int) {
	if l == r {
		this.tree[treeIndex] = this.data[l]
		return
	}
	leftTreeIndex := leftChild(treeIndex)
	rightTreeIndex := rightChild(treeIndex)

	mid := l + (r-l)/2
	this.buildSegmentTree(leftTreeIndex, l, mid)
	this.buildSegmentTree(rightTreeIndex, mid+1, r)

	this.tree[treeIndex] = this.merger(this.tree[leftTreeIndex], this.tree[rightTreeIndex])
}

func (this *SegmentTree) GetSize() int {
	return len(this.data)
}

func (this *SegmentTree) Get(index int) interface{} {
	if index < 0 || index >= len(this.data) {
		panic("Index is illegal.")
	}

	return this.data[index]
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
func leftChild(index int) int {
	return index*2 + 1
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
func rightChild(index int) int {
	return index*2 + 2
}

// 返回区间[queryL, queryR]的值
func (this *SegmentTree) Query(queryL int, queryR int) interface{} {
	if queryL < 0 || queryL >= len(this.data) || queryR < 0 || queryR > len(this.data) || queryL > queryR {
		panic("Index is illegal.")
	}

	return this.query(0, 0, len(this.data)-1, queryL, queryR)
}

// 在以treeIndex为根的线段树中[l...r]的范围里，搜索区间[queryL...queryR]的值
func (this *SegmentTree) query(treeIndex int, l int, r int, queryL int, queryR int) interface{} {
	if l == queryL && r == queryR {
		return this.tree[treeIndex]
	}

	mid := l + (r-l)/2

	// treeIndex的节点分为[l...mid]和[mid+1...r]两部分
	leftTreeIndex := leftChild(treeIndex)
	rightTreeIndex := rightChild(treeIndex)

	if queryL >= mid+1 {
		return this.query(rightTreeIndex, mid+1, r, queryL, queryR)
	} else if queryR <= mid {
		return this.query(leftTreeIndex, l, mid, queryL, queryR)
	}

	leftResult := this.query(leftTreeIndex, l, mid, queryL, mid)
	rightResult := this.query(rightTreeIndex, mid+1, r, mid+1, queryR)
	return this.merger(leftResult, rightResult)
}

func (this *SegmentTree) String() string {
	buffer := bytes.Buffer{}

	buffer.WriteString("[")
	for i := 0; i < len(this.tree); i++ {
		if this.tree[i] != nil {
			buffer.WriteString(fmt.Sprint(this.tree[i]))
		} else {
			buffer.WriteString("nil")
		}

		if i != len(this.tree)-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")

	return buffer.String()
}
