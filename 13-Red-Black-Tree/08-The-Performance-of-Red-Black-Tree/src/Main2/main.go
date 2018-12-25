package main

import (
	"math/rand"
	"math"
	"time"
	"Play-with-Data-Structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/src/BSTMap"
	"fmt"
	"Play-with-Data-Structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/src/AVLTree"
	"Play-with-Data-Structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/src/RBTree"
)

func main()  {
	n := 20000000

	var testData []int
	for i := 0; i < n; i++ {
		testData = append(testData, rand.Intn(math.MaxInt64))
	}

	// Test BST
	startTime := time.Now()

	bst := BSTMap.Constructor()
	for _, v := range testData {
		bst.Add(v, nil)
	}
	fmt.Println("BST: ", time.Now().Sub(startTime))

	// Test AVL
	startTime = time.Now()
	avl := AVLTree.Constructor()
	for _, v := range testData {
		avl.Add(v, nil)
	}
	fmt.Println("AVL: ", time.Now().Sub(startTime))

	// Test RBTree
	rbt := RBTree.Constructor()
	for _, v := range testData {
		rbt.Add(v, nil)
	}
	fmt.Println("RBTree: ", time.Now().Sub(startTime))
}
