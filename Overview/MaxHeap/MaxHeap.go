package MaxHeap

import (
	"Play-with-Data-Structures/Overview/Array"
	"Play-with-Data-Structures/Utils/Interfaces"
)

// 最大堆
type MaxHeap struct {
	data *Array.Array
}

func Constructor(capacity int) *MaxHeap {
	return &MaxHeap{
		data: Array.Constructor(capacity),
	}
}

func GetMaxHeapFromArr(arr []interface{}) *MaxHeap {
	a := Array.Constructor(len(arr))
	for _, v := range arr {
		a.AddLast(v)
	}

	maxHeap := &MaxHeap{a}
	for i := parent(len(arr) - 1); i >= 0; i-- {
		maxHeap.SiftDown(i)
	}
	return maxHeap
}

func leftChild(key int) int {
	return 2*key + 1
}

func rightChild(key int) int {
	return 2*key + 2
}

func parent(key int) int {
	if key == 0 {
		panic("Root don't have parent.")
	}

	return (key - 1) / 2
}

func (h *MaxHeap) GetSize() int {
	return h.data.GetSize()
}

func (h *MaxHeap) IsEmpty() bool {
	return h.data.IsEmpty()
}

func (h *MaxHeap) Add(e interface{}) {
	h.data.AddLast(e)
	h.SiftUp(h.data.GetSize() - 1)
}

func (h *MaxHeap) SiftUp(index int) {
	for index > 0 && Interfaces.Compare(h.data.Get(index), h.data.Get(parent(index))) > 0 {
		parentIndex := parent(index)
		// 交换父子元素
		temp := h.data.Get(index)
		h.data.Set(index, h.data.Get(parentIndex))
		h.data.Set(parentIndex, temp)
		// 更新 index 为父节点
		index = parentIndex
	}
}

func (h *MaxHeap) FindMax() interface{} {
	if h.data.GetSize() == 0 {
		panic("Empty maxHeap.")
	}
	return h.data.Get(0)
}

// 最大堆删除的都是堆顶元素
func (h *MaxHeap) Remove() interface{} {
	e := h.FindMax()
	h.data.Set(0, h.data.Get(h.data.GetSize()-1))
	h.data.RemoveLast()
	h.SiftDown(0)

	return e
}

func (h *MaxHeap) SiftDown(index int) {
	// 外部循环条件：当前节点不是叶子节点，即左孩子索引如果>=元素总数，说明已经数组越界
	for leftChild(index) < h.data.GetSize() {
		l := leftChild(index)
		r := l + 1

		// 判断与左右子节点哪个交换
		swapIndex := l
		if r < h.data.GetSize() && Interfaces.Compare(h.data.Get(r), h.data.Get(l)) > 0 {
			swapIndex = r
		}
		// 父节点大于较大的子节点，循环结束
		if Interfaces.Compare(h.data.Get(index), h.data.Get(l)) > 0 {
			break
		}
		// 交换父子值
		temp := h.data.Get(swapIndex)
		h.data.Set(swapIndex, h.data.Get(index))
		h.data.Set(index, temp)
		// 更新 index
		index = swapIndex
	}
}

// 替换堆的最大元素，并返回被替换的元素值
func (h *MaxHeap) Replace(e interface{}) interface{} {
	max := h.FindMax()
	h.data.Set(0, e)
	h.SiftDown(0)

	return max
}
