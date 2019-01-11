package main

import (
	"Play-with-Data-Structures/08-Heap-and-Priority-Queue/05-Heapify-and-Replace-in-Heap/src/MaxHeap"
	"Play-with-Data-Structures/Utils/Interfaces"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func testHeap(testData []interface{}, isHeapify bool) time.Duration {
	startTime := time.Now()

	var maxHeap *MaxHeap.MaxHeap
	if isHeapify {
		maxHeap = MaxHeap.GetMaxHeapFromArr(testData)
	} else {
		maxHeap = MaxHeap.Constructor()
		for _, v := range testData {
			maxHeap.Add(v)
		}
	}

	arr := make([]interface{}, len(testData))
	for i := 0; i < len(testData); i++ {
		arr[i] = maxHeap.ExtractMax()
	}
	for i := 1; i < len(testData); i++ {
		if Interfaces.Compare(arr[i-1], arr[i]) < 0 {
			panic("Error")
		}
	}

	fmt.Println("Test MaxHeap completed.")

	return time.Now().Sub(startTime)
}

func main() {
	n := 1000000

	var testData []interface{}
	rand.Seed(time.Now().Unix())
	for i := 0; i < n; i++ {
		testData = append(testData, rand.Intn(math.MaxInt32))
	}

	time1 := testHeap(testData, false)
	fmt.Println("Without heapify: ", time1)

	time2 := testHeap(testData, true)
	fmt.Println("With heapify: ", time2)
}
