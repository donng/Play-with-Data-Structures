package PriorityQueue

import "Play-with-Data-Structures/08-Heap-and-Priority-Queue/06-Priority-Queue/src/MaxHeap"

type PriorityQueue struct {
	maxHeap *MaxHeap.MaxHeap
}

func (q *PriorityQueue) GetSize() int {
	return q.maxHeap.Size()
}

func (q *PriorityQueue) IsEmpty() bool {
	return q.maxHeap.IsEmpty()
}

func (q *PriorityQueue) Enqueue(e interface{}) {
	q.maxHeap.Add(e)
}

func (q *PriorityQueue) Dequeue() interface{} {
	return q.maxHeap.ExtractMax()
}

func (q *PriorityQueue) GetFront() interface{} {
	return q.maxHeap.FindMax()
}

