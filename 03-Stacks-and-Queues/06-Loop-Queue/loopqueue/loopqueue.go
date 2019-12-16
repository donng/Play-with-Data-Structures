package loopqueue

type LoopQueue struct {
	data  []interface{}
	front int
	tail  int
	size  int
}

func New(capacity int) (lq *LoopQueue) {
	lq = &LoopQueue{}
	// 队列存在以空位，所以参数为20就需要开辟21的位置
	lq.data = make([]interface{}, capacity+1)
	lq.front = 0
	lq.tail = 0
	lq.size = 0

	return
}

func (lq *LoopQueue) GetCapacity() int {
	// 与上对应，需要删除空位
	return len(lq.data) - 1
}

func (lq *LoopQueue) GetSize() int {
	return lq.size
}

func (lq *LoopQueue) IsEmpty() bool {
	return lq.front == lq.tail
}

func (lq *LoopQueue) Enqueue(interface{}) {
	panic("implement me")
}

func (lq *LoopQueue) Eequeue() interface{} {
	panic("implement me")
}

func (lq *LoopQueue) GetFront() interface{} {
	panic("implement me")
}
