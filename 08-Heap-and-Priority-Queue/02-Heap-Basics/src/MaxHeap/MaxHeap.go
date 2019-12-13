package MaxHeap

import "Play-with-Data-Structures/08-Heap-and-Priority-Queue/02-Heap-Basics/src/Array"

type MaxHeap struct {
	data *Array.Array
}

func Constructor() *MaxHeap {
	return &MaxHeap{
		data: Array.Constructor(20),
	}
}

// 返回堆中的元素个数
func (h *MaxHeap) Size() int {
	return h.data.GetSize()
}

// 返回一个布尔值, 表示堆中是否为空
func (h *MaxHeap) IsEmpty() bool {
	return h.data.IsEmpty()
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的父亲节点的索引
func parent(index int) int {
	if index == 0 {
		panic("index-0 doesn't have parent.")
	}
	return (index - 1) / 2
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的左孩子节点的索引
func leftChild(index int) int {
	return index*2 + 1
}

// 返回完全二叉树的数组表示中，一个索引所表示的元素的右孩子节点的索引
func rightChild(index int) int {
	return index*2 + 2
}
