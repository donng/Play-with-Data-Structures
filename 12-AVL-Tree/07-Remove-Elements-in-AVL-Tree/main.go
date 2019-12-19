package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/donng/Play-with-Data-Structures/12-AVL-Tree/07-Remove-Elements-in-AVL-Tree/AVLTree"
	"github.com/donng/Play-with-Data-Structures/12-AVL-Tree/07-Remove-Elements-in-AVL-Tree/BSTMap"
	"github.com/donng/Play-with-Data-Structures/utils"
)

func main() {
	fmt.Println("Pride and Prejudice")

	filename, _ := filepath.Abs("pride-and-prejudice.txt")
	words := utils.ReadFile(filename)
	fmt.Println("Total words: ", len(words))

	// Test BST
	startTime := time.Now()

	bst := BSTMap.New()
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

	diffTime := time.Now().Sub(startTime)
	fmt.Println("BST: ", diffTime)

	// Test AVL Tree
	startTime = time.Now()

	avl := AVLTree.New()
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

	diffTime = time.Now().Sub(startTime)
	fmt.Println("AVL: ", diffTime)
}
