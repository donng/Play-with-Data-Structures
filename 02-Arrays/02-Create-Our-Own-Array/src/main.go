package main

import (
	"Play-with-Data-Structures/02-Arrays/02-Create-Our-Own-Array/src/Array"
	"fmt"
)

func main() {
	a := Array.GetArray(5)

	fmt.Println(a)
	fmt.Println(a.GetCapacity(), a.GetSize(), a.IsEmpty())
}
