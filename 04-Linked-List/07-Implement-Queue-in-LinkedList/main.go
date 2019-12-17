package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/donng/Play-with-Data-Structures/04-Linked-List/07-Implement-Queue-in-LinkedList/arrayqueue"
	"github.com/donng/Play-with-Data-Structures/04-Linked-List/07-Implement-Queue-in-LinkedList/linkedlistqueue"
	"github.com/donng/Play-with-Data-Structures/04-Linked-List/07-Implement-Queue-in-LinkedList/loopqueue"
	"github.com/donng/Play-with-Data-Structures/04-Linked-List/07-Implement-Queue-in-LinkedList/queue"
)

// 测试使用queue运行opCount个enqueueu和dequeue操作所需要的时间，单位：秒
func testQueue(queue queue.Queue, opCount int) float64 {
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

	arrayQueue := arrayqueue.New(20)
	time := testQueue(arrayQueue, opCount)
	fmt.Println("ArrayQueue, time:", time, "s")

	loopQueue := loopqueue.New(20)
	time = testQueue(loopQueue, opCount)
	fmt.Println("LoopQueue, time:", time, "s")

	linkedListQueue := linkedlistqueue.New()
	time = testQueue(linkedListQueue, opCount)
	fmt.Println("linkedListQueue, time:", time, "s")
}
