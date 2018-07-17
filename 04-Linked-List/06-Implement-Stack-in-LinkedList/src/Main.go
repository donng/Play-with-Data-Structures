package main

import (
	"fmt"
	"time"
	"math/rand"
)

func testStack(stack Stack, opCount int) float64 {
	startTime := time.Now()

	for i := 0; i < opCount; i++ {
		stack.Push(rand.Int())
	}
	for i := 0; i < opCount; i++ {
		stack.Pop()
	}

	return time.Now().Sub(startTime).Seconds()
}

func main() {
	//stack := GetLinkedListStack()
	//
	//for i := 0; i < 5; i++ {
	//	stack.Push(i)
	//	fmt.Println(stack)
	//}
	//
	//stack.Pop()
	//fmt.Println(stack)
	opCount := 10000000

	arrayStack := GetArrayStack(20)
	time1 := testStack(arrayStack, opCount)
	fmt.Println("ArrayStack, time: ", time1, " s")

	linkedListStack := GetLinkedListStack()
	time2 := testStack(linkedListStack, opCount)
	fmt.Println("ArrayStack, time: ", time2, " s")

	// 其实这个时间比较很复杂，因为 LinkedListStack 中包含更多的 new 操作
}
