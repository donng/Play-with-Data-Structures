package main

import "fmt"

func main() {
	a := GetArray(5)

	fmt.Println(a)
	fmt.Println(a.GetCapacity(), a.GetSize(), a.IsEmpty())
}
