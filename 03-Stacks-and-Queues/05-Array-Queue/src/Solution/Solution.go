package main

/// Leetcode 102. Binary Tree Level Order Traversal
/// https://leetcode.com/problems/binary-tree-level-order-traversal/description/
/// 二叉树的层序遍历
///
/// 二叉树的层序遍历是一个典型的可以借助队列解决的问题。
/// 该代码主要用于使用Leetcode上的问题测试我们的ArrayQueue。
/// 对于二叉树的层序遍历，这个课程后续会讲到。
/// 届时，同学们也可以再回头看这个代码。
/// 不过到时，大家应该已经学会自己编写二叉树的层序遍历啦：）
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
func (this *Array) GetCapacity() int {
	return len(this.data)
}

// 获得数组中的元素个数
func (this *Array) GetSize() int {
	return this.size
}

// 返回数组是否为空
func (this *Array) IsEmpty() bool {
	return this.size == 0
}

// 在第 index 个位置插入一个新元素 e
func (this *Array) Add(index int, e interface{}) {
	if index < 0 || index > this.GetCapacity() {
		panic("Add failed. Require index >= 0 and index <= size.")
	}

	if this.size == len(this.data) {
		this.resize(2 * this.size)
	}

	for i := this.size - 1; i >= index; i-- {
		this.data[i+1] = this.data[i]
	}

	this.data[index] = e
	this.size++
}

// 向所有元素后添加一个新元素
func (this *Array) AddLast(e interface{}) {
	this.Add(this.size, e)
}

// 向所有元素前添加一个新元素
func (this *Array) AddFirst(e interface{}) {
	this.Add(0, e)
}

// 获取 index 索引位置的元素
func (this *Array) Get(index int) interface{} {
	if index < 0 || index >= this.size {
		panic("Get failed. Index is illegal.")
	}
	return this.data[index]
}

func (this *Array) GetLast() interface{} {
	return this.Get(this.size - 1)
}

func (this *Array) GetFirst() interface{} {
	return this.Get(0)
}

// 修改 index 索引位置的元素
func (this *Array) Set(index int, e interface{}) {
	if index < 0 || index >= this.size {
		panic("Set failed. Index is illegal.")
	}
	this.data[index] = e
}

// 查找数组中是否有元素 e
func (this *Array) Contains(e interface{}) bool {
	for i := 0; i < this.size; i++ {
		if this.data[i] == e {
			return true
		}
	}
	return false
}

// 查找数组中元素 e 所在的索引，不存在则返回 -1
func (this *Array) Find(e interface{}) int {
	for i := 0; i < this.size; i++ {
		if this.data[i] == e {
			return i
		}
	}
	return -1
}

// 查找数组中元素 e 所有的索引组成的切片，不存在则返回 -1
func (this *Array) FindAll(e interface{}) (indexes []int) {
	for i := 0; i < this.size; i++ {
		if this.data[i] == e {
			indexes = append(indexes, i)
		}
	}
	return
}

// 从数组中删除 index 位置的元素，返回删除的元素
func (this *Array) Remove(index int) interface{} {
	if index < 0 || index >= this.size {
		panic("Remove failed,Index is illegal.")
	}

	e := this.data[index]
	for i := index + 1; i < this.size; i++ {
		this.data[i-1] = this.data[i]
	}
	this.size--
	this.data[this.size] = nil //loitering object != memory leak

	// 考虑边界条件，避免长度为 1 时，resize 为 0
	if this.size == len(this.data)/4 && len(this.data)/2 != 0 {
		this.resize(len(this.data) / 2)
	}
	return e
}

// 从数组中删除第一个元素，返回删除的元素
func (this *Array) RemoveFirst() interface{} {
	return this.Remove(0)
}

// 从数组中删除最后一个元素，返回删除的元素
func (this *Array) RemoveLast() interface{} {
	return this.Remove(this.size - 1)
}

// 从数组中删除一个元素 e
func (this *Array) RemoveElement(e interface{}) {
	index := this.Find(e)
	if index != -1 {
		this.Remove(index)
	}
}

// 从数组中删除所有元素 e
func (this *Array) RemoveAllElement(e interface{}) bool {
	if this.Find(e) == -1 {
		return false
	}

	for i := 0; i < this.size; i++ {
		if this.data[i] == e {
			this.Remove(i)
		}
	}
	return true
}

// 为数组扩容
func (this *Array) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity)
	for i := 0; i < this.size; i++ {
		newData[i] = this.data[i]
	}

	this.data = newData
}

type ArrayQueue struct {
	array *Array
}

func GetArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		array: GetArray(capacity),
	}
}

func (this *ArrayQueue) GetSize() int {
	return this.array.GetSize()
}

func (this *ArrayQueue) IsEmpty() bool {
	return this.array.IsEmpty()
}

func (this *ArrayQueue) GetCapacity() int {
	return this.array.GetCapacity()
}

func (this *ArrayQueue) Enqueue(e interface{}) {
	this.array.AddLast(e)
}

func (this *ArrayQueue) Dequeue() interface{} {
	return this.array.RemoveFirst()
}

func (this *ArrayQueue) GetFront() interface{} {
	return this.array.GetFirst()
}

type Queue interface {
	GetSize() int
	IsEmpty() bool
	Enqueue(interface{})
	Dequeue() interface{}
	GetFront() interface{}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
