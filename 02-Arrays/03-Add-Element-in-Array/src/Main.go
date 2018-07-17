package main

import "fmt"

func main() {
	a := GetArray(5)

	a.AddLast(2)
	a.AddLast(3)
	a.AddFirst(1)
	fmt.Println(a)
}
