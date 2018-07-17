package main

import (
	"time"
	"math/rand"
	"fmt"
)

func testQueue(queue Queue, opCount int) float64 {
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

	arrayQueue := GetArrayQueue(20)
	time := testQueue(arrayQueue, opCount)
	fmt.Println("ArrayQueue, time:", time, "s")

	loopQueue := GetLoopQueue(20)
	time = testQueue(loopQueue, opCount)
	fmt.Println("LoopQueue, time:", time, "s")

	/**
	 * 测试结果：
	 * ArrayQueue, time: 7.94818435 s
	 * LoopQueue, time: 0.020557388 s
	 */
}
