package main

import (
	"Play-with-Data-Structures/03-Stacks-and-Queues/02-Array-Stack/src/ArrayStack"
	"fmt"
)

func main() {
	stack := ArrayStack.GetArrayStack(10)
	fmt.Println(stack)
	for i := 0; i < 10; i++ {
		stack.Push(i)
		fmt.Println(stack)
	}

	stack.Pop()
	fmt.Println(stack)
}
