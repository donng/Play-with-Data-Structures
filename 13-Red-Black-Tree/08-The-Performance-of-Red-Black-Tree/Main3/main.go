package main

import (
	"fmt"
	"time"

	"github.com/donng/Play-with-Data-Structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/AVLTree"
	"github.com/donng/Play-with-Data-Structures/13-Red-Black-Tree/08-The-Performance-of-Red-Black-Tree/RBTree"
)

func main() {
	n := 20000000

	var testData []int
	for i := 0; i < n; i++ {
		testData = append(testData, i)
	}

	startTime := time.Now()
	avl := AVLTree.New()
	for _, v := range testData {
		avl.Add(v, nil)
	}
	fmt.Println("AVL: ", time.Now().Sub(startTime))

	startTime = time.Now()
	rbt := RBTree.New()
	for _, v := range testData {
		rbt.Add(v, nil)
	}
	fmt.Println("RBTree: ", time.Now().Sub(startTime))
}
