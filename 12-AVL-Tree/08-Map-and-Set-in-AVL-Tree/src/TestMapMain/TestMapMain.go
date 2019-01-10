package main

import (
	"Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/src/AVLTree"
	"Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/src/BSTMap"
	"Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/src/FileOperation"
	"Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/src/LinkedListMap"
	"Play-with-Data-Structures/12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/src/Map"
	"fmt"
	"path/filepath"
	"time"
)

func testSet(p Map.Map, filename string) time.Duration {
	startTime := time.Now()

	words := FileOperation.ReadFile(filename)
	fmt.Println("Total words:", len(words))

	for _, word := range words {
		if p.Contains(word) {
			p.Set(word, p.Get(word).(int)+1)
		} else {
			p.Add(word, 1)
		}
	}
	fmt.Println("Total different words: ", p.GetSize())
	fmt.Println("Frequency of PRIDE:", p.Get("pride"))
	fmt.Println("Frequency of PREJUDICE: ", p.Get("prejudice"))

	return time.Now().Sub(startTime)
}

func main() {
	filename, _ := filepath.Abs("12-AVL-Tree/08-Map-and-Set-in-AVL-Tree/pride-and-prejudice.txt")

	bstMap := BSTMap.Constructor()
	time1 := testSet(bstMap, filename)
	fmt.Println("BST Map:", time1)

	linkedListMap := LinkedListMap.Constructor()
	time2 := testSet(linkedListMap, filename)
	fmt.Println("linked List Map:", time2)

	avlMap := AVLTree.Constructor()
	time3 := testSet(avlMap, filename)
	fmt.Println("AVL Map:", time3)
}
