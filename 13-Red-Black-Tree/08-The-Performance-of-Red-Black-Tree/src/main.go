package main

import (
	"time"
	"fmt"
	"path/filepath"
	"Play-with-Data-Structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/src/FileOperation"
	"Play-with-Data-Structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/src/BSTMap"
	"Play-with-Data-Structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/src/AVLTree"
	"Play-with-Data-Structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/src/RBTree"
)

func main() {
	fmt.Println("Pride and Prejudice")

	filename, _ := filepath.Abs("13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/pride-and-prejudice.txt")
	words := FileOperation.ReadFile(filename)
	fmt.Println("Total words: ", len(words))

	// Test BST
	startTime := time.Now()
	bst := BSTMap.Constructor()
	for _, word := range words {
		if bst.Contains(word) {
			bst.Set(word, bst.Get(word).(int)+1)
		} else {
			bst.Add(word, 1)
		}
	}
	for _, word := range words {
		bst.Contains(word)
	}
	fmt.Println("BST: ", time.Now().Sub(startTime))

	// Test AVL
	startTime = time.Now()
	avl := AVLTree.Constructor()
	for _, word := range words {
		if avl.Contains(word) {
			avl.Set(word, avl.Get(word).(int)+1)
		} else {
			avl.Add(word, 1)
		}
	}
	for _, word := range words {
		avl.Contains(word)
	}
	fmt.Println("AVL: ", time.Now().Sub(startTime))

	// Test RBTree
	startTime = time.Now()
	rbt := RBTree.Constructor()
	for _, word := range words {
		if rbt.Contains(word) {
			rbt.Set(word, rbt.Get(word).(int)+1)
		} else {
			rbt.Add(word, 1)
		}
	}
	for _, word := range words {
		rbt.Contains(word)
	}
	fmt.Println("RBT: ", time.Now().Sub(startTime))
}
