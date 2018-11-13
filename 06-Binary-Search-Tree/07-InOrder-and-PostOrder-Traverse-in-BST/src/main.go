package main

import (
	"Play-with-Data-Structures/06-Binary-Search-Tree/07-InOrder-and-PostOrder-Traverse-in-BST/src/BST"
	"fmt"
)

func main() {
	bst := BST.GetBST()
	nums := [...]int{5, 3, 6, 8, 4, 2}
	for _, num := range nums {
		bst.Add(num)
	}

	/////////////////
	//      5      //
	//    /   \    //
	//   3    6    //
	//  / \    \   //
	// 2  4     8  //
	/////////////////
	bst.PreOrder()
	fmt.Println()

	bst.InOrder()
	fmt.Println()

	bst.PostOrder()
	// fmt.Println(bst)
}
