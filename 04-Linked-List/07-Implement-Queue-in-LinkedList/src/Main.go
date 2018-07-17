package main

import (
	"time"
	"fmt"
	"math/rand"
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

	linkedListpQueue := &LinkedListQueue{}
	time = testQueue(linkedListpQueue, opCount)
	fmt.Println("linkedListpQueue, time:", time, "s")

}
