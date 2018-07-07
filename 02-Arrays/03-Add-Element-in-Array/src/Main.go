package main

import "fmt"

func main() {
	a := getArray(5)

	a.addLast(2)
	a.addLast(3)
	a.addFirst(1)
	fmt.Println(a)
}
