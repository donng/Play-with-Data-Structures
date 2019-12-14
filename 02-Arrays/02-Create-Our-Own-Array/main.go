package main

import (
	"fmt"
	"github.com/donng/Play-with-Data-Structures/02-Arrays/02-Create-Our-Own-Array/array"
)

func main() {
	a := array.New(5)

	fmt.Println(a)
	fmt.Println(a.GetCapacity(), a.GetSize(), a.IsEmpty())
}
