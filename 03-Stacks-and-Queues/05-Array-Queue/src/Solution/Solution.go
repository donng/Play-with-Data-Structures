package main

import "Play-with-Data-Structures/Utils/Interfaces"

/// Leetcode 102. Binary Tree Level Order Traversal
/// https://leetcode.com/problems/binary-tree-level-order-traversal/description/
/// 二叉树的层序遍历
///
/// 二叉树的层序遍历是一个典型的可以借助队列解决的问题。
/// 该代码主要用于使用Leetcode上的问题测试我们的ArrayQueue。
/// 对于二叉树的层序遍历，这个课程后续会讲到。
/// 届时，同学们也可以再回头看这个代码。
/// 不过到时，大家应该已经学会自己编写二叉树的层序遍历啦：）

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 将以下所有代码提交到leetcode即可通过

type Array struct {
	data []interface{}
	size int
}

// 构造函数，传入数组的容量capacity构造Array
func GetArray(capacity int) *Array {
	return &Array{
		data: make([]interface{}, capacity),
	}
}

// 获取数组的容量
func (a *Array) GetCapacity() int {
	return len(a.data)
}

// 获得数组中的元素个数
func (a *Array) GetSize() int {
	return a.size
}

// 返回数组是否为空
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

// 在第 index 个位置插入一个新元素 e
func (a *Array) Add(index int, e interface{}) {
	if index < 0 || index > a.size {
		panic("Add failed. Require index >= 0 and index <= size.")
	}

	if a.size == len(a.data) {
		a.resize(2 * a.size)
	}

	for i := a.size - 1; i >= index; i-- {
		a.data[i+1] = a.data[i]
	}

	a.data[index] = e
	a.size++
}

// 向所有元素后添加一个新元素
func (a *Array) AddLast(e interface{}) {
	a.Add(a.size, e)
}

// 向所有元素前添加一个新元素
func (a *Array) AddFirst(e interface{}) {
	a.Add(0, e)
}

// 获取 index 索引位置的元素
func (a *Array) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("Get failed. Index is illegal.")
	}
	return a.data[index]
}

func (a *Array) GetLast() interface{} {
	return a.Get(a.size - 1)
}

func (a *Array) GetFirst() interface{} {
	return a.Get(0)
}

// 修改 index 索引位置的元素
func (a *Array) Set(index int, e interface{}) {
	if index < 0 || index >= a.size {
		panic("Set failed. Index is illegal.")
	}
	a.data[index] = e
}

// 查找数组中是否有元素 e
func (a *Array) Contains(e interface{}) bool {
	for i := 0; i < a.size; i++ {
		if Interfaces.Compare(a.data[i], e) == 0 {
			return true
		}
	}
	return false
}

// 查找数组中元素 e 所在的索引，不存在则返回 -1
func (a *Array) Find(e interface{}) int {
	for i := 0; i < a.size; i++ {
		if Interfaces.Compare(a.data[i], e) == 0 {
			return i
		}
	}
	return -1
}

// 查找数组中元素 e 所有的索引组成的切片，不存在则返回 -1
func (a *Array) FindAll(e interface{}) (indexes []int) {
	for i := 0; i < a.size; i++ {
		if Interfaces.Compare(a.data[i], e) == 0 {
			indexes = append(indexes, i)
		}
	}
	return
}

// 从数组中删除 index 位置的元素，返回删除的元素
func (a *Array) Remove(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("Remove failed,Index is illegal.")
	}

	e := a.data[index]
	for i := index + 1; i < a.size; i++ {
		a.data[i-1] = a.data[i]
	}
	a.size--
	a.data[a.size] = nil //loitering object != memory leak

	// 考虑边界条件，避免长度为 1 时，resize 为 0
	if a.size == len(a.data)/4 && len(a.data)/2 != 0 {
		a.resize(len(a.data) / 2)
	}
	return e
}

// 从数组中删除第一个元素，返回删除的元素
func (a *Array) RemoveFirst() interface{} {
	return a.Remove(0)
}

// 从数组中删除最后一个元素，返回删除的元素
func (a *Array) RemoveLast() interface{} {
	return a.Remove(a.size - 1)
}

// 从数组中删除一个元素 e
func (a *Array) RemoveElement(e interface{}) {
	index := a.Find(e)
	if index != -1 {
		a.Remove(index)
	}
}

// 从数组中删除所有元素 e
func (a *Array) RemoveAllElement(e interface{}) bool {
	if a.Find(e) == -1 {
		return false
	}

	for i := 0; i < a.size; i++ {
		if Interfaces.Compare(a.data[i], e) == 0 {
			a.Remove(i)
		}
	}
	return true
}

// 为数组扩容
func (a *Array) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity)
	for i := 0; i < a.size; i++ {
		newData[i] = a.data[i]
	}

	a.data = newData
}

type ArrayQueue struct {
	array *Array
}

func GetArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		array: GetArray(capacity),
	}
}

func (a *ArrayQueue) GetSize() int {
	return a.array.GetSize()
}

func (a *ArrayQueue) IsEmpty() bool {
	return a.array.IsEmpty()
}

func (a *ArrayQueue) GetCapacity() int {
	return a.array.GetCapacity()
}

func (a *ArrayQueue) Enqueue(e interface{}) {
	a.array.AddLast(e)
}

func (a *ArrayQueue) Dequeue() interface{} {
	return a.array.RemoveFirst()
}

func (a *ArrayQueue) GetFront() interface{} {
	return a.array.GetFirst()
}

type Queue interface {
	GetSize() int
	IsEmpty() bool
	Enqueue(interface{})
	Dequeue() interface{}
	GetFront() interface{}
}

type Pair struct {
	key   *TreeNode
	value int
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	queue := GetArrayQueue(10)
	queue.Enqueue(Pair{root, 0})
	for !queue.IsEmpty() {
		front := queue.Dequeue().(Pair)
		node := front.key
		level := front.value

		if level == len(res) {
			res = append(res, []int{})
		}

		res[level] = append(res[level], node.Val)
		if node.Left != nil {
			queue.Enqueue(Pair{node.Left, level + 1})
		}
		if node.Right != nil {
			queue.Enqueue(Pair{node.Right, level + 1})
		}
	}

	return res
}
