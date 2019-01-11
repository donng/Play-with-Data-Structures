package main

import (
	"Play-with-Data-Structures/06-Binary-Search-Tree/11-Remove-Min-and-Max-in-BST/src/BST"
	"Play-with-Data-Structures/Utils/Interfaces"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	bst := BST.Constructor()

	n := 1000
	var nums []interface{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		bst.Add(rand.Intn(10000))
	}

	for !bst.IsEmpty() {
		nums = append(nums, bst.RemoveMin())
	}
	fmt.Println(nums)

	for i := 1; i < len(nums); i++ {
		if Interfaces.Compare(nums[i-1], nums[i]) > 0 {
			panic("Error")
		}
	}
	fmt.Println("removeMin test completed.")
}
