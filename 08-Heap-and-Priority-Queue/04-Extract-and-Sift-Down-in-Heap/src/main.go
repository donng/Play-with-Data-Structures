package main

import (
	"Play-with-Data-Structures/08-Heap-and-Priority-Queue/04-Extract-and-Sift-Down-in-Heap/src/MaxHeap"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	n := 1000000

	maxHeap := MaxHeap.Constructor()
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		randNum := rand.Intn(math.MaxInt32)
		maxHeap.Add(randNum)
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = maxHeap.ExtractMax().(int)
	}

	for i := 1; i < n; i++ {
		if arr[i-1] < arr[i] {
			panic("Error")
		}
	}

	fmt.Println("Test MaxHeap completed.")
}
