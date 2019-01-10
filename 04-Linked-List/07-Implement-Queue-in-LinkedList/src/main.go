package main

import (
	"Play-with-Data-Structures/04-Linked-List/07-Implement-Queue-in-LinkedList/src/ArrayQueue"
	"Play-with-Data-Structures/04-Linked-List/07-Implement-Queue-in-LinkedList/src/LinkedListQueue"
	"Play-with-Data-Structures/04-Linked-List/07-Implement-Queue-in-LinkedList/src/LoopQueue"
	"Play-with-Data-Structures/04-Linked-List/07-Implement-Queue-in-LinkedList/src/Queue"
	"fmt"
	"math/rand"
	"time"
)

// 测试使用queue运行opCount个enqueueu和dequeue操作所需要的时间，单位：秒
func testQueue(queue Queue.Queue, opCount int) float64 {
	startTime := time.Now()

	for i := 0; i < opCount; i++ {
		queue.Enqueue(rand.Int())
	}
	for i := 0; i < opCount; i++ {
		queue.Dequeue()
	}

	return time.Now().Sub(startTime).Seconds()
}

func main() {
	opCount := 100000

	arrayQueue := ArrayQueue.Constructor(20)
	time := testQueue(arrayQueue, opCount)
	fmt.Println("ArrayQueue, time:", time, "s")

	loopQueue := LoopQueue.Constructor(20)
	time = testQueue(loopQueue, opCount)
	fmt.Println("LoopQueue, time:", time, "s")

	linkedListQueue := LinkedListQueue.Constructor()
	time = testQueue(linkedListQueue, opCount)
	fmt.Println("linkedListQueue, time:", time, "s")
}
