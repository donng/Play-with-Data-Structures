package main

import (
	"Play-with-Data-Structures/02-Arrays/04-Query-and-Update-Element/src/Array"
	"fmt"
)

func main() {
	arr := Array.Constructor(20)
	for i := 0; i < 10; i++ {
		arr.AddLast(i)
	}
	fmt.Println(arr)

	arr.Add(1, 100)
	fmt.Println(arr)

	arr.AddFirst(-1)
	fmt.Println(arr)
	// [-1, 0, 100, 1, 2, 3, 4, 5, 6, 7, 8, 9]
}
