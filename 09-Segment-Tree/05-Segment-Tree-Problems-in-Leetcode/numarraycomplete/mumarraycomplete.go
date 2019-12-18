package main

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
		panic("index is out of range")
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

// 返回区间[queryL, queryR]的值
func (t *SegmentTree) Query(queryL int, queryR int) interface{} {
	if queryL < 0 || queryL >= len(t.data) || queryR < 0 || queryR > len(t.data) {
		panic("index is out of range")
	}

	return t.query(0, 0, len(t.data)-1, queryL, queryR)
}

// 在以treeIndex为根的线段树中[l...r]的范围里，搜索区间[queryL...queryR]的值
func (t *SegmentTree) query(treeIndex int, l int, r int, queryL int, queryR int) interface{} {
	if l == queryL && r == queryR {
		return t.tree[treeIndex]
	}

	mid := l + (r-l)/2

	// treeIndex的节点分为[l...mid]和[mid+1...r]两部分
	leftTreeIndex := leftChild(treeIndex)
	rightTreeIndex := rightChild(treeIndex)
	if queryL >= mid+1 {
		return t.query(rightTreeIndex, mid+1, r, queryL, queryR)
	} else if queryR <= mid {
		return t.query(leftTreeIndex, l, mid, queryL, queryR)
	}

	leftResult := t.query(leftTreeIndex, l, mid, queryL, mid)
	rightResult := t.query(rightTreeIndex, mid+1, r, mid+1, queryR)
	return t.merger(leftResult, rightResult)
}

func (t *SegmentTree) String() string {
	buffer := bytes.Buffer{}

	buffer.WriteString("[")
	for i := 0; i < len(t.tree); i++ {
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

type NumArray struct {
	segmentTree *SegmentTree
}

func New(nums []int) NumArray {
	data := make([]interface{}, len(nums))
	for i := 0; i < len(nums); i++ {
		data[i] = nums[i]
	}

	setTree := GetSegmentTree(data, func(i interface{}, j interface{}) interface{} {
		return i.(int) + j.(int)
	})

	return NumArray{setTree}
}

func (this *NumArray) SumRange(i int, j int) int {
	if this.segmentTree == nil {
		panic("Segment Tree is null")
	}
	return this.segmentTree.Query(i, j).(int)
}

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	obj := New(nums)

	fmt.Println(obj.SumRange(0, 2))
	fmt.Println(obj.SumRange(2, 5))
	fmt.Println(obj.SumRange(0, 5))
}
