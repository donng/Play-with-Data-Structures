package main

import (
	"fmt"

	"github.com/donng/Play-with-Data-Structures/02-Arrays/03-Add-Element-in-Array/array"
)

func main() {
	a := array.New(5)

	a.AddLast(2)
	a.AddLast(3)
	a.AddFirst(1)

	fmt.Println(a)
}
