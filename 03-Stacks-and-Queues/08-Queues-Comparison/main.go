package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/donng/Play-with-Data-Structures/03-Stacks-and-Queues/08-Queues-Comparison/arrayqueue"
	"github.com/donng/Play-with-Data-Structures/03-Stacks-and-Queues/08-Queues-Comparison/loopqueue"
	"github.com/donng/Play-with-Data-Structures/03-Stacks-and-Queues/08-Queues-Comparison/queue"
)

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
	time1 := testQueue(arrayQueue, opCount)
	fmt.Println("ArrayQueue, time:", time1, "s")

	loopQueue := loopqueue.New(20)
	time2 := testQueue(loopQueue, opCount)
	fmt.Println("LoopQueue, time:", time2, "s")

	/**
	 * 测试结果：
	 * ArrayQueue, time: 7.94818435 s
	 * LoopQueue, time: 0.020557388 s
	 */
}
