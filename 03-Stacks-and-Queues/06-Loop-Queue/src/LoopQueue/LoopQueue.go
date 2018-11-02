package LoopQueue

type LoopQueue struct {
	data  []interface{}
	front int
	tail  int
	size  int
}

func GetLoopQueue(capacity int) (l *LoopQueue) {
	l = &LoopQueue{}
	// 队列存在以空位，所以参数为20就需要开辟21的位置
	l.data = make([]interface{}, capacity+1)
	l.front = 0
	l.tail = 0
	l.size = 0

	return
}

func (l *LoopQueue) GetCapacity() int {
	// 与上对应，需要删除空位
	return len(l.data) - 1
}

func (l *LoopQueue) GetSize() int {
	return l.size
}

func (l *LoopQueue) IsEmpty() bool {
	return l.front == l.tail
}

func (l *LoopQueue) Enqueue(interface{}) {
	panic("implement me")
}

func (l *LoopQueue) Eequeue() interface{} {
	panic("implement me")
}

func (l *LoopQueue) GetFront() interface{} {
	panic("implement me")
}
