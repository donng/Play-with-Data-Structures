package main

import (
	"fmt"
	"github.com/donng/Play-with-Data-Structures/03-Stacks-and-Queues/02-Array-Stack/arraystack"
)

func main() {
	stack := arraystack.New(10)
	fmt.Println(stack)
	for i := 0; i < 10; i++ {
		stack.Push(i)
		fmt.Println(stack)
	}

	stack.Pop()
	fmt.Println(stack)
}
