package main

import (
	"Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/src/AVLSet"
	"Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/src/BSTSet"
	"Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/src/LinkedListSet"
	"Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/src/Set"
	"Play-with-Data-Structures/Utils/FileOperation"
	"fmt"
	"path/filepath"
	"time"
)

func testSet(set Set.Set, filename string) time.Duration {
	startTime := time.Now()

	words := FileOperation.ReadFile(filename)
	fmt.Println("Total words:", len(words))
	for _, word := range words {
		set.Add(word)
	}
	fmt.Println("Total different words:", set.GetSize())

	return time.Now().Sub(startTime)
}

func main() {
	filename, _ := filepath.Abs("../../pride-and-prejudice.txt")

	bstSet := BSTSet.Constructor()
	time1 := testSet(bstSet, filename)
	fmt.Println("BST Set :", time1)

	linkedListSet := LinkedListSet.Constructor()
	time2 := testSet(linkedListSet, filename)
	fmt.Println("linked List Set:", time2)

	avlSet := AVLSet.Constructor()
	time3 := testSet(avlSet, filename)
	fmt.Println("AVL Set:", time3)
}
