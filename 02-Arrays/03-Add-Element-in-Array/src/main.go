package main

import (
	"Play-with-Data-Structures/02-Arrays/03-Add-Element-in-Array/src/Array"
	"fmt"
)

func main() {
	a := Array.Constructor(5)

	a.AddLast(2)
	a.AddLast(3)
	a.AddFirst(1)
	fmt.Println(a)
}
