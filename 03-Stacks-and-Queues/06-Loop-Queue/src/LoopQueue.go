package main

type LoopQueue struct {
	data  []interface{}
	front int
	tail  int
	size  int
}

func GetLoopQueue(capacity int) (l *LoopQueue) {
	l = &LoopQueue{}
	l.data = make([]interface{}, capacity)
	l.front = 0
	l.tail = 0
	l.size = 0

	return
}

func (l *LoopQueue) GetCapacity() int {
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
