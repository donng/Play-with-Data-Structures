package main

import (
	"fmt"
	"Play-with-Data-Structures/06-Binary-Search-Tree/06-PreOrder-Traverse-in-BST/src/BST"
)

func main() {
	bst := BST.GetBST()
	nums := [...]int{5, 3, 6, 8, 4, 2}
	for _, num := range nums {
		bst.Add(num)
	}
	bst.PreOrder()

	fmt.Println(bst)
}
