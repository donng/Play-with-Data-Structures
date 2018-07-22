package main

import "fmt"

func main() {
	bst := &BST{}
	nums := [...]int{5, 3, 6, 8, 4, 2}
	for _, num := range nums {
		bst.Add(num)
	}
	bst.PreOrder()

	fmt.Println(bst)
}
