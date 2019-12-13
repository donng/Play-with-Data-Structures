package Solution

/// Leetcode 102. Binary Tree Level Order Traversal
/// https://leetcode.com/problems/binary-tree-level-order-traversal/description/
/// 二叉树的层序遍历
///
/// 二叉树的层序遍历是一个典型的可以借助队列解决的问题。
/// 该代码主要用于使用Leetcode上的问题测试我们的LoopQueue。
/// 对于二叉树的层序遍历，这个课程后续会讲到。
/// 届时，同学们也可以再回头看这个代码。
/// 不过到时，大家应该已经学会自己编写二叉树的层序遍历啦：）
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 将以下所有代码提交到leetcode即可通过

// 在这一版LoopQueue的实现中，我们将不浪费那1个空间：）
type Queue interface {
	GetSize() int
	IsEmpty() bool
	Enqueue(interface{})
	Dequeue() interface{}
	GetFront() interface{}
}

type LoopQueue struct {
	data  []interface{}
	front int
	tail  int
	size  int
}

func Constructor(capacity int) *LoopQueue {
	// 由于不浪费空间，所以data静态数组的大小是capacity
	// 而不是capacity + 1
	return &LoopQueue{
		data: make([]interface{}, capacity),
	}
}

func (lq *LoopQueue) GetCapacity() int {
	return len(lq.data)
}

func (lq *LoopQueue) GetSize() int {
	return lq.size
}

func (lq *LoopQueue) IsEmpty() bool {
	// 注意，我们不再使用front和tail之间的关系来判断队列是否为空，而直接使用size
	return lq.size == 0
}

// 入队
func (lq *LoopQueue) Enqueue(e interface{}) {
	// 注意，我们不再使用front和tail之间的关系来判断队列是否为满，而直接使用size
	capacity := lq.GetCapacity()
	if lq.size == capacity {
		lq.resize(capacity * 2)
	}
	lq.data[lq.tail] = e
	lq.tail = (lq.tail + 1) % len(lq.data)
	lq.size++
}

// 获得队列头部元素
func (lq *LoopQueue) Dequeue() (e interface{}) {
	if lq.IsEmpty() {
		panic("Cannot dequeue from empty queue")
	}

	e = lq.data[lq.front]
	lq.data[lq.front] = nil
	// 循环队列需要执行求余运算
	lq.front = (lq.front + 1) % len(lq.data)
	lq.size--
	if lq.size == lq.GetCapacity()/4 && lq.size != 0 {
		lq.resize(lq.GetCapacity() / 2)
	}

	return
}

// 查看队列头部元素
func (lq *LoopQueue) GetFront() interface{} {
	if lq.IsEmpty() {
		panic("Queue is empty")
	}

	return lq.data[lq.front]
}

func (lq *LoopQueue) resize(capacity int) {
	length := len(lq.data)
	newData := make([]interface{}, capacity)
	for i := 0; i < lq.size; i++ {
		newData[i] = lq.data[(i+lq.front)%length]
	}
	lq.data = newData
	lq.front = 0
	lq.tail = lq.size
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

	queue := Constructor(10)
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
