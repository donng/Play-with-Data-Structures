package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/donng/Play-with-Data-Structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/AVLTree"
	"github.com/donng/Play-with-Data-Structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/BSTMap"
	"github.com/donng/Play-with-Data-Structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/RBTree"
)

func main() {
	n := 20000000

	var testData []int
	for i := 0; i < n; i++ {
		testData = append(testData, rand.Intn(math.MaxInt64))
	}

	// Test BST
	startTime := time.Now()

	bst := BSTMap.New()
	for _, v := range testData {
		bst.Add(v, nil)
	}
	fmt.Println("BST: ", time.Now().Sub(startTime))

	// Test AVL
	startTime = time.Now()
	avl := AVLTree.New()
	for _, v := range testData {
		avl.Add(v, nil)
	}
	fmt.Println("AVL: ", time.Now().Sub(startTime))

	// Test RBTree
	rbt := RBTree.New()
	for _, v := range testData {
		rbt.Add(v, nil)
	}
	fmt.Println("RBTree: ", time.Now().Sub(startTime))
}
