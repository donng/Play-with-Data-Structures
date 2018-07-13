package main

type LoopQueue struct {
	data []interface{}
	front int
	tail int
	size int
}

func getLoopQueue(capacity int) (l *LoopQueue) {
	l = &LoopQueue{}
	l.data = make([]interface{}, capacity)
	l.front = 0
	l.tail = 0
	l.size = 0

	return
}

func (l *LoopQueue)getCapacity() int {
	return len(l.data) - 1
}

func (l *LoopQueue) getSize() int {
	return l.size
}

func (l *LoopQueue) isEmpty() bool {
	return l.front == l.tail
}

func (l *LoopQueue) enqueue(interface{}) {
	panic("implement me")
}

func (l *LoopQueue) dequeue() interface{} {
	panic("implement me")
}

func (l *LoopQueue) getFront() interface{} {
	panic("implement me")
}

