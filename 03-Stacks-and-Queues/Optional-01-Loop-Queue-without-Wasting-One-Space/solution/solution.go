package solution

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
}

func New(capacity int) *LoopQueue {
	return &LoopQueue{
		data: make([]interface{}, capacity+1),
	}
}

func (lq *LoopQueue) GetCapacity() int {
	return len(lq.data) - 1
}

func (lq *LoopQueue) GetSize() int {
	// 注意此时getSize的逻辑:
	// 如果tail >= front，非常简单，队列中的元素个数就是tail - front
	// 如果tail < front，说明我们的循环队列"循环"起来了，此时，队列中的元素个数为：
	// tail - front + data.length
	// 画画图，看能不能理解为什么？
	//
	// 也可以理解成，此时，data中没有元素的数目为front - tail,
	// 整体元素个数就是 data.length - (front - tail) = data.length + tail - front
	if lq.tail >= lq.front {
		return lq.tail - lq.front
	} else {
		return lq.tail - lq.front + len(lq.data)
	}
}

func (lq *LoopQueue) IsEmpty() bool {
	return lq.front == lq.tail
}

// 入队
func (lq *LoopQueue) Enqueue(e interface{}) {
	if (lq.tail+1)%len(lq.data) == lq.front {
		lq.resize(lq.GetCapacity() * 2)
	}
	lq.data[lq.tail] = e
	lq.tail = (lq.tail + 1) % len(lq.data)
}

// 获得队列头部元素
func (lq *LoopQueue) Dequeue() (e interface{}) {
	if lq.IsEmpty() {
		panic("cannot dequeue from empty queue")
	}

	e = lq.data[lq.front]
	lq.data[lq.front] = nil
	// 循环队列需要执行求余运算
	lq.front = (lq.front + 1) % len(lq.data)
	if lq.GetSize() == lq.GetCapacity()/4 && lq.GetSize() != 0 {
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
	newData := make([]interface{}, capacity+1)
	sz := lq.GetSize()
	for i := 0; i < sz; i++ {
		newData[i] = lq.data[(i+lq.front)%len(lq.data)]
	}
	lq.data = newData
	lq.front = 0
	lq.tail = sz
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

	queue := New(10)
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
