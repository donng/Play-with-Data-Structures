package main

import (
	"fmt"
	"Play-with-Data-Structures/03-Stacks-and-Queues/07-Implementation-of-Loop-Queue/src/LoopQueue"
)

func main() {
	queue := LoopQueue.Constructor(10)
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)

		if i%3 == 2 {
			queue.Dequeue()
			fmt.Println(queue)
		}
	}
}
