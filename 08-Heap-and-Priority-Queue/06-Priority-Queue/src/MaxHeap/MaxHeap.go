package MaxHeap

import (
	"Play-with-Data-Structures/08-Heap-and-Priority-Queue/05-Heapify-and-Replace-in-Heap/src/Array"
	"Play-with-Data-Structures/Utils/Interfaces"
)

type MaxHeap struct {
	data *Array.Array
}

func Constructor() *MaxHeap {
	return &MaxHeap{
		Array.Constructor(20),
	}
}

func GetMaxHeapFromArr(arr []interface{}) *MaxHeap {
	maxHeap := &MaxHeap{
		Array.GetArrayFromArr(arr),
	}

	for i := parent(len(arr) - 1); i >= 0; i-- {
		maxHeap.siftDown(i)
	}

	return maxHeap
}

// 返回堆中的元素个数
func (this *MaxHeap) Size() int {
	return this.data.GetSize()
}

// 返回一个布尔值, 表示堆中是否为空
func (this *MaxHeap) IsEmpty() bool {
	return this.data.IsEmpty()
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

// 向堆中添加元素
func (this *MaxHeap) Add(e interface{}) {
	this.data.AddLast(e)
	this.siftUp(this.data.GetSize() - 1)
}

func (this *MaxHeap) siftUp(k int) {
	for k > 0 && Interfaces.Compare(this.data.Get(k), this.data.Get(parent(k))) > 0 {
		this.data.Swap(k, parent(k))
		k = parent(k)
	}
}

func (this *MaxHeap) FindMax() interface{} {
	if this.data.GetSize() == 0 {
		panic("Can not findMax when heap is empty.")
	}
	return this.data.Get(0)
}

// 取出堆中最大元素
func (this *MaxHeap) ExtractMax() interface{} {
	ret := this.FindMax()

	this.data.Swap(0, this.data.GetSize()-1)
	this.data.RemoveLast()
	this.siftDown(0)

	return ret
}

// data[j] 是 leftChild 和 rightChild 中的最大值
func (this *MaxHeap) siftDown(k int) {
	for leftChild(k) < this.data.GetSize() {
		j := leftChild(k)
		// j+1是右孩子索引，如果存在右孩子比较后获得左右孩子中较大值的索引
		if j+1 < this.data.GetSize() && Interfaces.Compare(this.data.Get(j+1), this.data.Get(j)) > 0 {
			j++
		}
		// data[j] 是 leftChild 和 rightChild 中的最大值
		if Interfaces.Compare(this.data.Get(k), this.data.Get(j)) > 0 {
			break
		}

		this.data.Swap(k, j)
		k = j
	}
}

// 取出堆中的最大元素，并且替换成元素e
func (this *MaxHeap) Replace(e interface{}) interface{} {
	ret := this.FindMax()

	this.data.Set(0, e)
	this.siftDown(0)

	return ret
}
