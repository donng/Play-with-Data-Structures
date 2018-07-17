package main

import "fmt"

func main() {
	queue := GetArrayQueue(20)
	for i := 0; i < 10; i++ {
		queue.Enqueue(i)
		fmt.Println(queue)

		if i % 3 == 2 {
			queue.Dequeue()
			fmt.Println(queue)
		}
	}
}
