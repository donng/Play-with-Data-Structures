package main

import (
	"fmt"
	"math/rand"
	"time"
	"Play-with-Data-Structures/06-Binary-Search-Tree/11-Remove-Min-and-Max-in-BST/src/BST"
)

func main() {
	bst := BST.GetBST()

	n := 1000
	nums := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		bst.Add(rand.Intn(10000))
	}

	for !bst.IsEmpty() {
		nums = append(nums, bst.RemoveMin())
	}
	fmt.Println(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			panic("Error")
		}
	}
	fmt.Println("removeMin test completed.")
}
